package handler

import (
	"github.com/jmoiron/sqlx"
	api "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/apiserver/config"
	"github.com/modoki-paas/modoki-k8s/internal/grpcutil"
	"golang.org/x/xerrors"
)

type Plugin struct {
	Name    string
	Client  api.GeneratorClient
	Metrics bool
}

// ServerContext contains accessor used by handlers
type ServerContext struct {
	Config *config.Config

	DB        *sqlx.DB
	EnvConfig *config.EnvConfig

	AppClient     api.AppClient
	UserOrgClient api.UserOrgClient
	Generators    []*Plugin
}

func (sc *ServerContext) connectDB() error {
	d, err := sqlx.Open("mysql", sc.Config.DB)

	if err != nil {
		return xerrors.Errorf("failed to connect to database: %w", err)
	}
	sc.DB = d

	return nil
}

func (sc *ServerContext) connectAppClient() error {
	e := sc.Config.Endpoints.App

	conn, err := grpcutil.Dial(e.Endpoint, e.Insecure)

	if err != nil {
		return xerrors.Errorf("failed to dial grpc server: %w", err)
	}

	sc.AppClient = api.NewAppClient(conn)

	return nil
}

func (sc *ServerContext) connectGenerator(name string, endpoint string, insecure, metrics bool) error {
	conn, err := grpcutil.Dial(endpoint, insecure)

	if err != nil {
		return xerrors.Errorf("failed to dial grpc server: %w", err)
	}

	generator := api.NewGeneratorClient(conn)

	sc.Generators = append(sc.Generators, &Plugin{
		Name:    name,
		Client:  generator,
		Metrics: metrics,
	})

	return nil
}

func (sc *ServerContext) connectGenerators() error {
	e := sc.Config.Endpoints.Generator

	if err := sc.connectGenerator(
		"base",
		e.Endpoint,
		e.Insecure,
		true,
	); err != nil {
		return xerrors.Errorf("failed to connect to generator: %w", err)
	}

	for _, p := range sc.Config.Endpoints.Plugins {
		if err := sc.connectGenerator(
			p.Name,
			p.Endpoint.Endpoint,
			p.Endpoint.Insecure,
			p.MetricsAPI,
		); err != nil {
			return xerrors.Errorf("failed to connect to plugin(%s): %w", p.Name, err)
		}
	}

	return nil
}

func (sc *ServerContext) connectUserOrgClient() error {
	e := sc.Config.Endpoints.UserOrg

	conn, err := grpcutil.Dial(e.Endpoint, e.Insecure)

	if err != nil {
		return xerrors.Errorf("failed to dial grpc server: %w", err)
	}

	sc.UserOrgClient = api.NewUserOrgClient(conn)

	return nil
}

func NewServerContext(cfg *config.Config) (*ServerContext, error) {
	sctx := &ServerContext{}

	sctx.Config = cfg

	if err := sctx.connectDB(); err != nil {
		return nil, xerrors.Errorf("failed to connect to database: %w", err)
	}

	if err := sctx.connectAppClient(); err != nil {
		return nil, xerrors.Errorf("failed to connect to app server: %w", err)
	}

	if err := sctx.connectUserOrgClient(); err != nil {
		return nil, xerrors.Errorf("failed to connect to user_org server: %w", err)
	}

	if err := sctx.connectGenerators(); err != nil {
		return nil, xerrors.Errorf("failed to connect to generators server: %w", err)
	}

	return sctx, nil
}
