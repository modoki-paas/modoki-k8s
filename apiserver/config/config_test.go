package config

import (
	"io/ioutil"
	"os"
	"testing"

	"github.com/modoki-paas/modoki-k8s/internal/testutil"

	"github.com/google/go-cmp/cmp"
)

func TestReadDefaultWithEnv(t *testing.T) {
	rollback := testutil.SetTemporalEnv(
		"DB_USER=user",
		"DB_PASSWORD=password",
		"DB_HOST=host",
		"DB_PORT=2379",
		"DB_DATABASE=database",
		"MODOKI_API_KEY=apikey",
		"MODOKI_APP_DOMAIN=modoki.example.com",
	)

	defer rollback()

	cfg, err := ReadConfig()

	if err != nil {
		t.Fatalf("failed to load config: %+v", err)
	}

	expected := &Config{
		DB:        "user:password@tcp(host:2379)/database?parseTime=true",
		Address:   ":443",
		Namespace: "modoki",
		APIKeys: []string{
			"apikey",
		},
		Domain: "modoki.example.com",

		Endpoints: Endpoints{
			Generator: &Endpoint{
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

			Plugins: nil,
		},

		DBElements: dbElements{
			User:     "user",
			Password: "password",
			Host:     "host",
			Port:     "2379",
			Database: "database",
		},
	}

	if diff := cmp.Diff(expected, cfg); diff != "" {
		t.Errorf("the default config differs: %s", diff)
	}
}

func TestReadFile(t *testing.T) {
	cfgYAML := `
db: user:password@tcp(host:2379)/database?parseTime=true
domain: modoki.example.com
api_keys:
- apikey
namespace: modoki-app
address: :50001
endpoints:
  generator:
    endpoint: :50001
    insecure: false
  app:
    endpoint: :50001
    insecure: false
  user_org:
    endpoint: :50001
    insecure: true
  plugins:
  - name: mysql
    metrics_api: true
    endpoint: localhost:443
    insecure: true
  - name: redis
    metrics_api: false
    endpoint: localhost:80
    insecure: false
`
	ioutil.WriteFile("apiserver.yaml", []byte(cfgYAML), 0755)
	defer os.Remove("apiserver.yaml")

	cfg, err := ReadConfig()

	if err != nil {
		t.Fatalf("failed to load config: %+v", err)
	}

	expected := &Config{
		DB:        "user:password@tcp(host:2379)/database?parseTime=true",
		Address:   ":50001",
		Namespace: "modoki-app",
		APIKeys: []string{
			"apikey",
		},
		Domain: "modoki.example.com",

		Endpoints: Endpoints{
			Generator: &Endpoint{
				Endpoint: ":50001",
				Insecure: false,
			},
			UserOrg: &Endpoint{
				Endpoint: ":50001",
				Insecure: true,
			},
			App: &Endpoint{
				Endpoint: ":50001",
				Insecure: false,
			},

			Plugins: []Plugin{
				{
					Name:       "mysql",
					MetricsAPI: true,
					Endpoint: Endpoint{
						Endpoint: "localhost:443",
						Insecure: true,
					},
				},
				{
					Name:       "redis",
					MetricsAPI: false,
					Endpoint: Endpoint{
						Endpoint: "localhost:80",
						Insecure: false,
					},
				},
			},
		},
	}

	if diff := cmp.Diff(expected, cfg); diff != "" {
		t.Errorf("the default config differs: %s", diff)
	}
}
