// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package basic

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/golang/glog"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/auth"
	"github.com/pilotariak/trinquet/config"
	"github.com/pilotariak/trinquet/transport"
)

const (
	label = "BasicAuth"

	key = "basic"
)

type basicAuthSystem struct{}

func init() {
	auth.RegisterAuthentication(label, newBasicAuthSystem)
}

func newBasicAuthSystem(config *config.Configuration) (auth.Authentication, error) {
	return &basicAuthSystem{}, nil
}

func (ba basicAuthSystem) Name() string {
	return label
}

func (ba basicAuthSystem) Key() string {
	return key
}

func (ba basicAuthSystem) Credentials(ctx context.Context, parentSpan opentracing.Span, username string, password string) (string, error) {
	glog.V(2).Infof("Set credentials %s", username)
	auth := username + ":" + password
	token := base64.StdEncoding.EncodeToString([]byte(auth))
	return token, nil
}

func (ba basicAuthSystem) Authenticate(ctx context.Context, parentSpan opentracing.Span, token string) (map[string]string, error) {
	glog.V(2).Infof("Check BasicAuth token: %s", token)
	b, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return nil, fmt.Errorf("Can't check authentication: %s", err)
	}
	pair := strings.SplitN(string(b), ":", 2)
	if len(pair) != 2 {
		return nil, fmt.Errorf("Not Authorized")
	}
	// glog.V(2).Infof("Auth: %s / %s", pair[0], pair[1])
	if pair[0] != auth.Username || pair[1] != auth.Password {
		return nil, fmt.Errorf("Unauthorized")
	}

	headers := map[string]string{}
	headers[transport.Username] = pair[0]
	return headers, nil
}
