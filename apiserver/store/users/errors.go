package users

import "golang.org/x/xerrors"

var (
	// ErrUserIDDuplicates means user ids duplicate
	ErrUserIDDuplicates = xerrors.New("user ids duplicate")

	// ErrUnknownUser means user is unknown
	ErrUnknownUser = xerrors.New("unknown user")

	// ErrUnknownRoleBinding means role binding is unknown
	ErrUnknownRoleBinding = xerrors.New("unknown role binding")
)
