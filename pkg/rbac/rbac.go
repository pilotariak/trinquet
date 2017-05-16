// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

// Package rbac provides a lightweight role-based access control
package rbac

import (
	"fmt"

	"github.com/golang/glog"
)

const (
	AdminRole = "admin"

	UserRole = "user"

	GrpcHealthKey = "/grpc.health.v1.Health/Check"

	HealthRole = "health"
)

var userRights = map[string][]string{}

func Roles() map[string][]string {
	return userRights
}

func AddRoles(api string, service string, roles map[string][]string) {
	for name, role := range roles {
		userRights[fmt.Sprintf("/%s.%s/%s", api, service, name)] = role
	}
}

func HasRights(key string, roles []string) error {
	glog.V(2).Infof("Check Roles: %s %s %s", key, roles, userRights[key])
	if key == GrpcHealthKey && len(roles) == 1 && HealthRole == roles[0] {
		return nil
	}
	for _, value := range roles {
		for _, val := range userRights[key] {
			if val == value {
				return nil
			}
		}
	}
	return fmt.Errorf("Invalid user rights: %s", roles)
}
