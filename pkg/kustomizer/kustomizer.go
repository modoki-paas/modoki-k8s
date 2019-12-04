package kustomizer

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/otiai10/copy"
	"golang.org/x/xerrors"
	ktypes "sigs.k8s.io/kustomize/api/types"
	"sigs.k8s.io/yaml"
)

var (
	// OriginalDir is the path to the original yaml folder
	OriginalDir = "./templates"
)

type Workspace struct {
	Dir string
}

func NewWorkspace() (*Workspace, error) {
	d, err := ioutil.TempDir("", "yamler-")

	if err != nil {
		return nil, xerrors.Errorf("failed to create workspace dir: %w", err)
	}

	copy.Copy(OriginalDir, d)

	return &Workspace{Dir: d}, nil
}

func (w *Workspace) Close() error {
	return os.RemoveAll(w.Dir)
}

func (w *Workspace) CommandWithInput(ctx context.Context, input io.Reader, command string, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, command, args...)

	cmd.Dir = w.Dir
	cmd.Env = os.Environ()

	return cmd
}

func (w *Workspace) Command(ctx context.Context, command string, args ...string) *exec.Cmd {
	return w.CommandWithInput(ctx, bytes.NewReader(nil), command, args...)
}

func (w *Workspace) KustomizeWithInput(ctx context.Context, input io.Reader, args ...string) *exec.Cmd {
	return w.CommandWithInput(ctx, bytes.NewReader(nil), "kustomize", args...)
}

func (w *Workspace) Kustomize(ctx context.Context, args ...string) *exec.Cmd {
	return w.CommandWithInput(ctx, bytes.NewReader(nil), "kustomize", args...)
}

func (w *Workspace) Build(ctx context.Context) (string, error) {
	output, err := w.Kustomize(ctx, "build").CombinedOutput()

	if err != nil {
		return "", xerrors.Errorf("failed to execute kustomize build: %w", err)
	}

	return string(output), nil
}

func (w *Workspace) SaveConfig(y *ktypes.Kustomization) error {
	y.FixKustomizationPostUnmarshalling()
	if errs := y.EnforceFields(); len(errs) != 0 {
		return xerrors.Errorf("failed to enforce fields in kustomization yaml: %s", strings.Join(errs, ", "))
	}

	b, err := yaml.Marshal(y)

	if err != nil {
		return xerrors.Errorf("failed to marshal yaml for kustomize: %w", err)
	}

	if err := ioutil.WriteFile(filepath.Join(w.Dir, "kustomization.yaml"), b, 0774); err != nil {
		return xerrors.Errorf("failed to save yaml for kustomize: %w", err)
	}

	return nil
}

func (w *Workspace) LoadConfig() (*ktypes.Kustomization, error) {
	b, err := ioutil.ReadFile(filepath.Join(w.Dir, "kustomization.yaml"))

	if err != nil {
		return nil, xerrors.Errorf("failed to read yaml for kustomize: %w", err)
	}

	var y ktypes.Kustomization
	if err := yaml.Unmarshal(b, &y); err != nil {
		return nil, xerrors.Errorf("failed to parse yaml for kustomize: %w", err)
	}

	y.FixKustomizationPostUnmarshalling()
	if errs := y.EnforceFields(); len(errs) != 0 {
		return nil, xerrors.Errorf("failed to enforce fields in kustomization yaml: %s", strings.Join(errs, ", "))
	}

	return &y, nil
}
