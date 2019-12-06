package handler

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"io/ioutil"
	"testing"

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

	resp, err := h.Operate(context.Background(), &modoki.OperateRequest{
		Id:        "test-id",
		Kind:      modoki.OperateKind_Apply,
		Performer: 10,
		Spec: &modoki.AppSpec{
			Owner:   11,
			Name:    "app-name",
			Domain:  "app-name.example.com",
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

	if diff := cmp.Diff(string(b), resp.ApplyYaml.Config); diff != "" {
		t.Fatal(diff)
	}

	t.Log(resp.ApplyYaml.Config)
}
