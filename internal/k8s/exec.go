package k8s

import (
	"context"
	"golang.org/x/xerrors"
	"io"
	"os"
	"os/exec"
)

func (c *Client) newKubectlCommand(ctx context.Context, args ...string) *exec.Cmd {
	cmd := exec.CommandContext(ctx, c.kubectlPath, args...)

	cmd.Env = append(os.Environ(), "KUBECONFIG="+c.kubeconfig)

	return cmd
}

func (c *Client) Apply(ctx context.Context, yaml io.Reader) (string, error) {
	cmd := c.newKubectlCommand(ctx, "apply", "-f", "-")
	cmd.Stdin = yaml

	b, err := cmd.CombinedOutput()

	if err != nil {
		return string(b), xerrors.Errorf("failed to execute kubectl: %w", err)
	}

	return string(b), nil
}

func (c *Client) Delete(ctx context.Context, yaml io.Reader) (string, error) {
	cmd := c.newKubectlCommand(ctx, "delete", "-f", "-")
	cmd.Stdin = yaml

	b, err := cmd.CombinedOutput()

	if err != nil {
		return string(b), xerrors.Errorf("failed to execute kubectl: %w", err)
	}

	return string(b), nil
}
