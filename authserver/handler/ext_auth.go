package handler

import (
	"context"
	"strings"

	"github.com/google/martian/log"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"golang.org/x/xerrors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	extauth "github.com/envoyproxy/go-control-plane/envoy/service/auth/v2"
	envoy_type "github.com/envoyproxy/go-control-plane/envoy/type"
	"github.com/gogo/googleapis/google/rpc"
	rpcstatus "google.golang.org/genproto/googleapis/rpc/status"
)

type ExtAuthZ struct {
	GA      auth.GatewayAuthorizer
	Context *ServerContext
}

var _ extauth.AuthorizationServer = &ExtAuthZ{}

// Check handles requests from ext_authz in Envoy proxy
func (ea *ExtAuthZ) Check(ctx context.Context, req *extauth.CheckRequest) (*extauth.CheckResponse, error) {
	authzHeader := req.Attributes.Request.Http.Headers["authorization"]
	token := strings.TrimPrefix(authzHeader, "Bearer ")

	targetHeader := req.Attributes.Request.Http.Headers[strings.ToLower(auth.TargetIDHeader)]

	md, err := ea.GA.GetAuthenticatedMetadata(ctx, token, targetHeader)

	if xerrors.Is(err, auth.ErrUnauthenticated) {
		return &extauth.CheckResponse{
			Status: &rpcstatus.Status{
				Code: int32(rpc.UNAUTHENTICATED),
			},
			HttpResponse: &extauth.CheckResponse_DeniedResponse{
				DeniedResponse: &extauth.DeniedHttpResponse{
					Status: &envoy_type.HttpStatus{
						Code: envoy_type.StatusCode_Unauthorized,
					},
					Body: "UNAUTHENTICATED",
				},
			},
		}, nil

		return nil, status.Error(codes.Unauthenticated, "the token is missing or invalid")
	}

	if err != nil {
		log.Errorf("authentication failed: %+v", err)

		return &extauth.CheckResponse{
			Status: &rpcstatus.Status{
				Code: int32(rpc.INTERNAL),
			},
			HttpResponse: &extauth.CheckResponse_DeniedResponse{
				DeniedResponse: &extauth.DeniedHttpResponse{
					Status: &envoy_type.HttpStatus{
						Code: envoy_type.StatusCode_InternalServerError,
					},
					Body: "INTERNAL_SERVER_ERROR",
				},
			},
		}, nil
	}

	return &extauth.CheckResponse{
		Status: &rpcstatus.Status{
			Code: int32(rpc.OK),
		},
		HttpResponse: &extauth.CheckResponse_OkResponse{
			OkResponse: &extauth.OkHttpResponse{
				Headers: []*core.HeaderValueOption{
					{
						Header: &core.HeaderValue{
							Key:   strings.ToLower(auth.UserIDHeader),
							Value: md.UserID,
						},
					},
					{
						Header: &core.HeaderValue{
							Key:   strings.ToLower(auth.TargetIDHeader),
							Value: md.TargetID,
						},
					},
					{
						Header: &core.HeaderValue{
							Key:   strings.ToLower(auth.RolesHeader),
							Value: md.Roles.Marshal(),
						},
					},
					{
						Header: &core.HeaderValue{
							Key:   "authorization",
							Value: "Bearer " + ea.Context.Config.APIKeys[0],
						},
					},
				},
			},
		},
	}, nil
}
