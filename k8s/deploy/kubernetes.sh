#!/usr/bin/env bash

# Copyright (C) 2016-2019 Nicolas Lamirault <nicolas.lamirault@gmail.com>

# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

NO_COLOR="\033[0m"
OK_COLOR="\033[32;01m"
INFO_COLOR="\033[34;01m"
ERROR_COLOR="\033[31;01m"
WARN_COLOR="\033[33;01m"

MINIKUBE_CONFIG="./minikube/kube-config"
kubeconfig=""

clusters="minikube, gce"
environments="local, test or prod"

image_name="trinquet"
app="trinquet"

namespace=""

cluster=$1
env=$2
action=$3
image_tag=$4
build_number=$(date +%Y%m%d%H%M%S)


function usage {
    echo -e "${OK_COLOR}Usage${NO_COLOR} : $0 <cluster> <env> <action> <docker image version>"
    echo -e "${INFO_COLOR}Clusters${NO_COLOR} : ${clusters}"
    echo -e "${INFO_COLOR}Environments${NO_COLOR} : ${environments}"
    echo -e "${INFO_COLOR}Action${NO_COLOR} : create, destroy"
}


function kube_install {
    echo -e "${OK_COLOR}Downloading kubectl${NO_COLOR}"
    local kube_download="https://storage.googleapis.com/kubernetes-release/release"
    curl -LO ${kube_download}/$(curl -s ${kube_download}/stable.txt)/bin/linux/amd64/kubectl > ./kubectl
    chmod +x ./kubectl
    export PATH=${PATH}:.
    kubectl version --client
}


function kube_context {
    if  [ ! -x "$(command -v kubectl)" ]; then
        kube_install
    fi
    local name=$1
    local environment=$2
    echo  -e "${OK_COLOR}Cluster: ${name}"
    echo  -e "${OK_COLOR}Environment: ${environment}"
    if [ "local" = "$2" ]; then
        local context="atmos"
        namespace="-n atmos"
    else
        local context="${name}-atmos-${environment}"
    fi
    kube_config ${environment}
    echo -e "${OK_COLOR}Switch to Kubernetes context: ${context}${NO_COLOR}" >&2
    kubectl --kubeconfig=${kubeconfig} config use-context ${context} >&2 || exit 1
}


function kube_config {
    local environment=$1
    case ${environment} in
        local)
            kubeconfig=${MINIKUBE_CONFIG}
            ;;
        stg|dev|itg|prp|prod)
            kubeconfig=${CLUSTER_CONFIG}
            ;;
        *)
            echo -e "${ERROR_COLOR}Invalid environment: ${environment}${NO_COLOR}"
            exit 1
    esac
}


function kube_replace {
    environment=$1
    image_tag=$2
    build_number=$3
    dir=$4

    echo -e "${OK_COLOR}Generate Kubernetes files for:${NO_COLOR}"
    echo -e "Environment: ${environment}"
    # echo -e "Namespace: ${ns}"
    echo -e "App: ${app}"
    echo -e "Docker ${image_name}:${image_tag}"
    echo -e "Build: ${build_number}"
    rm -fr ${dir} && mkdir -p ${dir} && cp -r deploy/k8s/* ${dir}
    # find ${dir} -name "*.yaml" | xargs sed -i "s/__KUBE_NAMESPACE__/${ns}/g"
    find ${dir} -name "*.yaml" | xargs sed -i "s/__KUBE_APP__/${app}/g"
    find ${dir} -name "*.yaml" | xargs sed -i "s/__KUBE_COMMIT_ID__/${build_number}/g"
    find ${dir} -name "*.yaml" | xargs sed -i "s/__KUBE_ENV__/${environment}/g"
    if [ "local" = "${environment}" ]; then
        find ${dir} -name "*.yaml" | xargs sed -i "s#__CI_REGISTRY_IMAGE__#atmosdb#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__CI_REGISTRY_TAG__#${image_tag}#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__KUBE_IMAGE_POLICY__#Never#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__KUBE_NAME__#minikube#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__KUBE_DB_PASSWORD__#atmosdb#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__KUBE_USER_SERVICE_PASSWORD__#${ATMOSDB_USER_SERVICE_PASSWORD}#g"

    else
        docker_repo=${repo_name}
        # if [ "itg" = "${environment}" ]; then
        #     if [[ "${image_tag}" == *"-rc-"* ]]; then
        #         docker_repo=${repo_dev_name};
        #     fi
        # fi
        if [ "${branch}" != "master" ]; then
            docker_repo=${repo_dev_name};
        fi
        echo -e "Docker Registry: ${docker_repo}"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__CI_REGISTRY_IMAGE__#${private_registry}/${docker_repo}/${image_name}#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__CI_REGISTRY_TAG__#${image_tag}#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__KUBE_IMAGE_POLICY__#Always#g"
        find ${dir} -name "*.yaml" | xargs sed -i "s#__KUBE_NAME__#cloud#g"
    fi
}


function kube_directory {
    action=$1
    directory=$2
    if [ -d "${directory}" ]; then
        if [ -n "$(ls -A ${directory})" ]; then
            kubectl --kubeconfig=${kubeconfig} ${action} -f "${directory}" ${namespace}
        fi
    fi
}


function kube_files {
    local action=$1
    local directory=$2
    local environment=$3
    for file in $(ls ${directory}/*-${environment}.yaml 2>/dev/null); do
        kubectl --kubeconfig=${kubeconfig} ${action} -f "${file}" ${namespace}
    done
}


function kube_deploy {
    local environment=$1
    local dir=$2
    echo -e "${OK_COLOR}Current context: $(kubectl --kubeconfig=${kubeconfig} config current-context)${NO_COLOR}"
    kube_directory "apply" "${dir}/commons"
    kube_directory "apply" "${dir}/${cluster}/commons"
    kube_files "apply" "${dir}/${cluster}" ${environment}
}


function kube_undeploy {
    local environment=$1
    local dir=$2
    echo -e "${OK_COLOR}Current context: ${context}${NO_COLOR}"
    kube_directory "delete" "${dir}/commons"
    kube_directory "delete" "${dir}/${cluster}/commons"
    kube_files "delete" "${dir}/${cluster}" ${environment}
}


if [ $# -eq 0 ]; then
    usage
    exit 0
fi
if [ $# -ne 4 ]; then
    usage
    exit 1
fi


case ${cluster} in
    minikube)
        url=${minikube_url}
        ;;
    *)
        echo -e "${ERROR_COLOR}Invalid cluster: ${env}${NO_COLOR}"
        usage
        exit 1
        ;;
esac


case ${action} in
    create)
        kube_context ${cluster} ${env}
        dir="/tmp/${app}"
        kube_replace ${env} ${image_tag} ${build_number} ${dir}
        kube_deploy ${env} ${dir}
        ;;
    destroy)
        kube_context ${cluster} ${env}
        dir="/tmp/${app}"
        kube_replace ${env} ${image_tag} ${build_number} ${dir}
        kube_undeploy ${env} ${dir}
        ;;
    *)
        echo -e "${ERROR_COLOR}Invalid action: [${action}]${NO_COLOR}"
	    echo -e "Valid actions: create, destroy"
        exit 1
esac
