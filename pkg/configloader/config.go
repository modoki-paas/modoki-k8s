package configloader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/heetch/confita"
	"github.com/heetch/confita/backend"
	"github.com/heetch/confita/backend/env"
	"github.com/heetch/confita/backend/file"
	"golang.org/x/xerrors"
)

// ReadConfig is a common implementation to load config
func ReadConfig(ctx context.Context, name string, out interface{}) error {
	yamlName := fmt.Sprintf("%s.yaml", name)
	jsonName := fmt.Sprintf("%s.json", name)

	bs := []backend.Backend{
		env.NewBackend(),
		file.NewOptionalBackend(filepath.Join("/etc/modoki", yamlName)),
		file.NewOptionalBackend(filepath.Join("/etc/modoki", jsonName)),
	}
	if home, err := os.UserHomeDir(); err == nil {
		bs = append(bs, file.NewOptionalBackend(filepath.Join(home, ".config", yamlName)))
		bs = append(bs, file.NewOptionalBackend(filepath.Join(home, ".config", jsonName)))
	}
	if wd, err := os.Getwd(); err == nil {
		bs = append(bs, file.NewOptionalBackend(filepath.Join(wd, yamlName)))
		bs = append(bs, file.NewOptionalBackend(filepath.Join(wd, jsonName)))
	}

	loader := confita.NewLoader(bs...)

	err := loader.Load(ctx, out)

	if err != nil {
		return xerrors.Errorf("failed to load config: %w", err)
	}

	return nil
}
