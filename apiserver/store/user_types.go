package store

import (
	"database/sql/driver"

	"golang.org/x/xerrors"
)

type (
	// UserTypeEnum represents the type of users
	UserTypeEnum string

	// UserSystemRole represents the role of users
	UserSystemRole string
)

func (e *UserTypeEnum) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		*e = UserTypeEnum(v)
	case string:
		*e = UserTypeEnum(v)
	default:
		return xerrors.Errorf("failed to scan json for %v", v)
	}
	return nil
}

func (e UserTypeEnum) Value() (driver.Value, error) {
	return string(e), nil
}

func (e *UserSystemRole) Scan(src interface{}) error {
	switch v := src.(type) {
	case []byte:
		*e = UserSystemRole(v)
	case string:
		*e = UserSystemRole(v)
	default:
		return xerrors.Errorf("failed to scan json for %v", v)
	}
	return nil
}

func (e UserSystemRole) Value() (driver.Value, error) {
	return string(e), nil
}

var (
	// UserNormal means a individual user, not organization
	UserNormal UserTypeEnum = "user"
	// UserOrganization means an organization contains some users
	UserOrganization UserTypeEnum = "organization"

	// UserRoleAdmin means an admin user
	UserRoleAdmin UserSystemRole = "admin"
	// UserRoleNothing means a normal user
	UserRoleNothing UserSystemRole = "admin"
)
