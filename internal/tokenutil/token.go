package tokenutil

import (
	"crypto/rand"
	"encoding/base64"

	"golang.org/x/xerrors"
)

func GenerateRandomToken() (string, error) {
	var b [128]byte
	if _, err := rand.Read(b[:]); err != nil {
		return "", xerrors.Errorf("failed to generate secure token")
	}

	return base64.StdEncoding.EncodeToString(b[:]), nil
}
