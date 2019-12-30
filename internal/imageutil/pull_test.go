package imageutil_test

import (
	"testing"

	"github.com/modoki-paas/modoki-k8s/internal/imageutil"
)

func TestGetImageHash(t *testing.T) {
	t.Run("with tag", func(t *testing.T) {
		hash, err := imageutil.GetImageHash("modokipaas/modoki-k8s:2f050bb9972e151d54f463a15a6184a3c656b534")

		if err != nil {
			t.Fatalf("failed to get image hash: %+v", err)
		}

		expected := "1ca964f6ffb65f7eebe20d2e53737313221625adac28450308ff9962db4f5f56"
		if hash != expected {
			t.Errorf("hash is not correct(actual: %s, expected: %s)", hash, expected)
		}
	})

	t.Run("with sha256", func(t *testing.T) {
		hash, err := imageutil.GetImageHash("modokipaas/modoki-k8s@sha256:1ca964f6ffb65f7eebe20d2e53737313221625adac28450308ff9962db4f5f56")

		if err != nil {
			t.Fatalf("failed to get image hash: %+v", err)
		}

		expected := "1ca964f6ffb65f7eebe20d2e53737313221625adac28450308ff9962db4f5f56"
		if hash != expected {
			t.Errorf("hash is not correct(actual: %s, expected: %s)", hash, expected)
		}
	})
}
