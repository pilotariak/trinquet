package utils

import (
	"os"

	"github.com/golang/glog"
)

const (
	// UsernameEnvVar is the environment variable that points to the username
	UsernameEnvVar = "TRINQUET_USERNAME"

	// ApikeyEnvVar is the environment variable that points to the APIKEY
	ApikeyEnvVar = "TRINQUET_APIKEY"

	// GrpcAddr is the environment variable that points to the gRPC server
	GrpcAddr = "TRINQUET_SERVER"
)

var (
	Username string
	Password string

	ServerAddress string
	// RestAddress   string
)

func setupFromEnvironmentVariables() {
	Username = os.Getenv(UsernameEnvVar)
	Password = os.Getenv(ApikeyEnvVar)
	ServerAddress = os.Getenv(GrpcAddr)
	glog.V(2).Infof("Env: %s %s %s", Username, Password, ServerAddress)
}
