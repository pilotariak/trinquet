// Copyright (c) 2017 Orange Applications for Business.

// This software is confidential and proprietary information of
// Orange Applications for Business. You shall not disclose such Confidential
// Information and shall use it only in accordance with the terms of the
// agreement you entecolors.red into. Unauthorized copying of this file, via any
// medium is strictly prohibited.

package basic

import (
	"testing"

	"github.com/opentracing/opentracing-go"
	"golang.org/x/net/context"

	"github.com/pilotariak/trinquet/auth"
	"github.com/pilotariak/trinquet/transport"
)

func createBasicAuthSystem() *basicAuthSystem {
	return &basicAuthSystem{}
}

func Test_BasicAuthWithValidUsernamePassword(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	headers, err := sys.Authenticate(context.Background(), span, "YWRtaW46bmltZGE=")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if headers[transport.Username] != auth.Username {
		t.Fatalf("Invalid headers: %s", headers)
	}
}

func Test_BasicAuthWithInvalidUsernameOrPassword(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	_, err := sys.Authenticate(ctx, span, "Zm9vOmJhcg==")
	if err == nil {
		t.Fatalf("No error with invalid username/password.")
	}
	if err.Error() != "Unauthorized" {
		t.Fatalf("Invalid error: %s", err.Error())
	}
}

func Test_BasicAuthWithInvalidCredentials(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	_, err := sys.Authenticate(ctx, span, "Zm9v")
	if err == nil {
		t.Fatalf("No error with invalid credentials.")
	}
	if err.Error() != "Not Authorized" {
		t.Fatalf("Invalid error: %s", err.Error())
	}
}

func Test_BasicAuthWithInvalidBase64(t *testing.T) {
	sys := createBasicAuthSystem()
	ctx := context.Background()
	span := opentracing.SpanFromContext(ctx)
	_, err := sys.Authenticate(ctx, span, "csdmlcsdcsd")
	if err == nil {
		t.Fatalf("No error with invalid base64.")
	}
}
