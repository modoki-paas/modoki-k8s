package apps

import "golang.org/x/xerrors"

var (
	// ErrAppNameDuplicates means app names duplicate
	ErrAppNameDuplicates = xerrors.New("app names duplicate")
)
