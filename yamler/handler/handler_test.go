package handler

import (
	"context"
	"io/ioutil"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/pkg/kustomizer"
	"github.com/modoki-paas/modoki-k8s/yamler/config"
)

func Test_OperateApply(t *testing.T) {
	h := Handler{
		Config: &config.Config{
			AppSecretName: "secret-name",
		},
	}

	kustomizer.OriginalDir = "../templates"

	ctx := context.Background()

	ctx = auth.AddTargetIDContext(ctx, "owner-id")

	resp, err := h.Operate(ctx, &modoki.OperateRequest{
		Id:   "test-id",
		Domain: "app-name.example.com",
		Kind: modoki.OperateKind_Apply,
		Spec: &modoki.AppSpec{
			Image:   "image-name",
			Command: []string{"command1", "command2"},
			Args:    []string{"arg1", "arg2"},
			Env: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
		},
		K8SConfig: &modoki.KubernetesConfig{
			Namespace: "namespace-name",
		},
	})

	if err != nil {
		t.Fatalf("failed to generate yaml: %+v", err)
	}

	b, err := ioutil.ReadFile("./testdata/desired.yaml")

	if err != nil {
		t.Fatalf("unknown config: %+v", err)
	}

	if diff := cmp.Diff(string(b), resp.Yaml.Config); diff != "" {
		t.Fatal(diff)
	}

	t.Log(resp.Yaml.Config)
}
