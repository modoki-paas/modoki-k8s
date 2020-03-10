package handler

import (
	"context"

	auth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
)

type ExtAuthZ struct {
}

var _ auth.AuthorizationServer = &ExtAuthZ{}

// Check handles requests from ext_authz in Envoy proxy
func (ea *ExtAuthZ) Check(ctx context.Context, req *auth.CheckRequest) (*auth.CheckResponse, error) {
	// req.Attributes.Request.Http.Headers
}
