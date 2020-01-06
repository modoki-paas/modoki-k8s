package tokens

import "golang.org/x/xerrors"

var (
	// ErrTokenIDDuplicates means token names duplicate
	ErrTokenIDDuplicates = xerrors.New("token names duplicate")

	// ErrUnknownToken means unknown token
	ErrUnknownToken = xerrors.New("unknown token")
)
