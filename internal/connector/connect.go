package connector

import (
	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/internal/grpcutil"
	"golang.org/x/xerrors"
)

type Connector struct {
	dialer *grpcutil.GRPCDialer
}

func NewConnector(dialer *grpcutil.GRPCDialer) *Connector {
	return &Connector{
		dialer: dialer,
	}
}

func (c *Connector) ConnectAppClient(endpoint string, insecure bool) (modoki.AppClient, error) {
	conn, err := c.dialer.Dial(endpoint, insecure)

	if err != nil {
		return nil, xerrors.Errorf("failed to dial app server: %w", err)
	}

	return modoki.NewAppClient(conn), nil
}

func (c *Connector) ConnectGenerator(endpoint string, insecure bool) (modoki.GeneratorClient, error) {
	conn, err := c.dialer.Dial(endpoint, insecure)

	if err != nil {
		return nil, xerrors.Errorf("failed to dial generator server: %w", err)
	}

	return modoki.NewGeneratorClient(conn), nil
}

func (c *Connector) ConnectUserOrgClient(endpoint string, insecure bool) (modoki.UserOrgClient, error) {
	conn, err := c.dialer.Dial(endpoint, insecure)

	if err != nil {
		return nil, xerrors.Errorf("failed to dial user/org server: %w", err)
	}

	return modoki.NewUserOrgClient(conn), nil
}

func (c *Connector) ConnectTokenClient(endpoint string, insecure bool) (modoki.TokenClient, error) {
	conn, err := c.dialer.Dial(endpoint, insecure)

	if err != nil {
		return nil, xerrors.Errorf("failed to dial token server: %w", err)
	}

	return modoki.NewTokenClient(conn), nil
}
