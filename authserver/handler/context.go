package handler

import (
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/authserver/config"
	"github.com/modoki-paas/modoki-k8s/internal/connector"
	"github.com/modoki-paas/modoki-k8s/internal/grpcutil"
	"golang.org/x/xerrors"
)

// ServerContext contains accessor used by handlers
type ServerContext struct {
	Config *config.Config

	AppClient     api.AppClient
	UserOrgClient api.UserOrgClient
	TokenClient   api.TokenClient

	Connector *connector.Connector
}

func (sc *ServerContext) connectAppClient() error {
	e := sc.Config.Endpoints.App

	var err error
	sc.AppClient, err = sc.Connector.ConnectAppClient(e.Endpoint, e.Insecure)

	if err != nil {
		return xerrors.Errorf("failed to dial app server: %w", err)
	}

	return nil
}

func (sc *ServerContext) connectUserOrgClient() error {
	e := sc.Config.Endpoints.App

	var err error
	sc.UserOrgClient, err = sc.Connector.ConnectUserOrgClient(e.Endpoint, e.Insecure)

	if err != nil {
		return xerrors.Errorf("failed to dial user/org server: %w", err)
	}

	return nil
}

func (sc *ServerContext) connectTokenClient() error {
	e := sc.Config.Endpoints.UserOrg

	var err error
	sc.TokenClient, err = sc.Connector.ConnectTokenClient(e.Endpoint, e.Insecure)

	if err != nil {
		return xerrors.Errorf("failed to dial token server: %w", err)
	}

	return nil
}

func NewServerContext(cfg *config.Config) (*ServerContext, error) {
	sctx := &ServerContext{}

	sctx.Config = cfg

	// TODO: api key for dialer
	sctx.Connector = connector.NewConnector(grpcutil.NewGRPCDialer(cfg.APIKeys[0]))

	if err := sctx.connectAppClient(); err != nil {
		return nil, xerrors.Errorf("failed to connect to app server: %w", err)
	}

	if err := sctx.connectUserOrgClient(); err != nil {
		return nil, xerrors.Errorf("failed to connect to user_org server: %w", err)
	}

	if err := sctx.connectTokenClient(); err != nil {
		return nil, xerrors.Errorf("failed to connect to generators server: %w", err)
	}

	return sctx, nil
}
