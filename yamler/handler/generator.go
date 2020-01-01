package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"sort"

	modoki "github.com/modoki-paas/modoki-k8s/api"
	"github.com/modoki-paas/modoki-k8s/internal/imageutil"
	"github.com/modoki-paas/modoki-k8s/pkg/auth"
	"github.com/modoki-paas/modoki-k8s/pkg/kustomizer"
	"github.com/modoki-paas/modoki-k8s/yamler/config"
	"golang.org/x/xerrors"
	"sigs.k8s.io/kustomize/api/resid"
	"sigs.k8s.io/kustomize/api/types"
)

type yamlerKustomizer func(ctx context.Context, cfg *config.Config, y *types.Kustomization, req *modoki.OperateRequest) (*types.Kustomization, error)

func setupNamespace(ctx context.Context, cfg *config.Config, y *types.Kustomization, req *modoki.OperateRequest) (*types.Kustomization, error) {
	y.Namespace = req.K8SConfig.Namespace

	return y, nil
}

func setupName(ctx context.Context, cfg *config.Config, y *types.Kustomization, req *modoki.OperateRequest) (*types.Kustomization, error) {
	y.NameSuffix = fmt.Sprintf("-%s", req.Id)

	return y, nil
}

func setupLabels(ctx context.Context, cfg *config.Config, y *types.Kustomization, req *modoki.OperateRequest) (*types.Kustomization, error) {
	y.CommonLabels = map[string]string{
		"modoki.tsuzu.xyz/id":     req.Id,
		"modoki.tsuzu.xyz/owner":  auth.GetTargetIDContext(ctx),
		"modoki.tsuzu.xyz/domain": req.Domain,
	}

	return y, nil
}

func setupIngress(ctx context.Context, cfg *config.Config, y *types.Kustomization, req *modoki.OperateRequest) (*types.Kustomization, error) {
	ingPatches, err := json.Marshal(kustomizer.Patches{
		{
			Op:    kustomizer.OpReplace,
			Path:  "/spec/tls/0/hosts/0",
			Value: req.Domain,
		},
		{
			Op:    kustomizer.OpReplace,
			Path:  "/spec/rules/0/host",
			Value: req.Domain,
		},
		{
			Op:    kustomizer.OpReplace,
			Path:  "/spec/tls/0/secretName",
			Value: cfg.AppSecretName,
		},
	})

	if err != nil {
		return nil, xerrors.Errorf("failed to encode patches for ingress: %w", err)
	}

	y.PatchesJson6902 = append(y.PatchesJson6902, types.PatchJson6902{
		Target: &types.PatchTarget{
			Gvk: resid.Gvk{
				Group:   "extensions",
				Version: "v1beta1",
				Kind:    "Ingress",
			},
			Name: "modoki-app-ing",
		},
		Patch: string(ingPatches),
	})

	return y, nil
}

func setupPod(ctx context.Context, cfg *config.Config, y *types.Kustomization, req *modoki.OperateRequest) (*types.Kustomization, error) {
	ow, err := imageutil.ParseOverwrite(req.Spec.Image, true)

	if err != nil {
		return nil, xerrors.Errorf("invalid image name parameter: %w", err)
	}

	if ow.Tag == "" && ow.Digest == "" {
		ow.Tag = "latest"
	}

	y.Images = []types.Image{
		{
			Name:    "IMAGE_NAME",
			NewName: ow.Name,
			NewTag:  ow.Tag,
			Digest:  ow.Digest,
		},
	}

	type Env struct {
		Name  string `json:"name" yaml:"name"`
		Value string `json:"value" yaml:"value"`
	}

	envs := make([]Env, 0, len(req.Spec.Env))
	for k, v := range req.Spec.Env {
		envs = append(envs, Env{
			Name:  k,
			Value: v,
		})
	}

	sort.Slice(envs, func(i int, j int) bool {
		return envs[i].Name < envs[j].Name
	})

	patches := kustomizer.Patches{}

	if len(req.Spec.Command) != 0 {
		patches = append(patches, kustomizer.Patch{
			Op:    kustomizer.OpAdd,
			Path:  "/spec/template/spec/containers/0/command",
			Value: req.Spec.Command,
		})
	}

	if len(req.Spec.Args) != 0 {
		patches = append(patches, kustomizer.Patch{
			Op:    kustomizer.OpAdd,
			Path:  "/spec/template/spec/containers/0/args",
			Value: req.Spec.Args,
		})
	}

	patches = append(patches, kustomizer.Patch{
		Op:    kustomizer.OpAdd,
		Path:  "/spec/template/spec/containers/0/env",
		Value: envs,
	})

	ingPatches, err := json.Marshal(patches)

	if err != nil {
		return nil, xerrors.Errorf("failed to encode patches for deployment: %w", err)
	}

	y.PatchesJson6902 = append(y.PatchesJson6902, types.PatchJson6902{
		Target: &types.PatchTarget{
			Gvk: resid.Gvk{
				Group:   "apps",
				Version: "v1",
				Kind:    "Deployment",
			},
			Name: "modoki-app-deploy",
		},
		Patch: string(ingPatches),
	})

	return y, nil
}
