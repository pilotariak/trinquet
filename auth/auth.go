// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package auth

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/config"
)

const (
	Username string = "admin"
	Password string = "nimda"
)

type AuthenticationFunc func(config *config.Configuration) (Authentication, error)

var registeredSystems = map[string](AuthenticationFunc){}

func RegisterAuthentication(name string, f AuthenticationFunc) {
	registeredSystems[name] = f
}

// Authentication define an authentication system
type Authentication interface {

	// Name identify the system
	Name() string

	// Scheme used into the :authorization header
	Key() string

	// Credentials check username and password and returns a token
	Credentials(ctx context.Context, parentSpan opentracing.Span, username string, password string) (string, error)

	// Authenticate check the authentication challenge
	Authenticate(ctx context.Context, parentSpan opentracing.Span, token string) (map[string]string, error)
}

// New returns a new authentication system using the name
func New(conf *config.Configuration) (Authentication, error) {
	glog.V(1).Infof("Authentication setup: %s", conf.Auth)
	if conf.Auth == nil {
		return nil, fmt.Errorf("Invalid authentication configuration: %s", conf)
	}
	f, ok := registeredSystems[conf.Auth.Name]
	if !ok {
		return nil, fmt.Errorf("Unsupported authentication system: %s", conf.Auth.Name)
	}
	system, err := f(conf)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func GetAuthenticationHeader(authentication Authentication, token string) string {
	return fmt.Sprintf("%s %s", authentication.Key(), token)
}
