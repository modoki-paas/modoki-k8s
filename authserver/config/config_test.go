package config

import (
	"testing"

	"github.com/modoki-paas/modoki-k8s/internal/testutil"

	"github.com/google/go-cmp/cmp"
)

func TestReadDefaultWithEnv(t *testing.T) {
	rollback := testutil.SetTemporalEnv(
		"OIDC_CLIENT_ID=clientid",
		"OIDC_CLIENT_SECRET=clientsecret",
		"OIDC_SCOPES=openid,profile",
		"OIDC_REDIRECT_URL=redirecturl",
		"OIDC_PROVIDER_URL=providerurl",
		"MODOKI_API_KEY=apikey",
	)

	defer rollback()

	cfg, err := ReadConfig()

	if err != nil {
		t.Fatalf("failed to load config: %+v", err)
	}

	expected := &Config{
		Address: ":443",
		APIKeys: []string{
			"apikey",
		},

		Endpoints: Endpoints{
			Token: &Endpoint{
				Endpoint: ":443",
				Insecure: true,
			},
			UserOrg: &Endpoint{
				Endpoint: ":443",
				Insecure: true,
			},
			App: &Endpoint{
				Endpoint: ":443",
				Insecure: true,
			},
		},
		OIDC: OpenIDConnect{
			ClientID:     "clientid",
			ClientSecret: "clientsecret",
			Scopes:       []string{"openid", "profile"},
			RedirectURL:  "redirecturl",
			ProviderURL:  "providerurl",
		},
	}

	if diff := cmp.Diff(expected, cfg); diff != "" {
		t.Errorf("the default config differs: %s", diff)
	}
}
