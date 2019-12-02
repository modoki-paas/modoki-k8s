package main

import (
	"bytes"
	"context"
	"io"
	"io/ioutil"
	"os"
	"os/exec"

	"github.com/docker/docker/pkg/testutil/cmd"
	"github.com/otiai10/copy"

	"golang.org/x/xerrors"
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

func (w *Workspace) CommandWithInput(ctx context.Context, input io.Reader, command string, args ...string) *cmd.Cmd {
	cmd := exec.CommandContext(ctx, command, args...)

	cmd.Dir = w.Dir
	cmd.Env = os.Environ()

	return cmd
}

func (w *Workspace) Command(ctx context.Context, command string, args ...string) *exec.Cmd {
	return w.CommandWithInput(ctx, bytes.NewReader(nil), command, args...)
}
