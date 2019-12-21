package permissions

var (
	// OrgCreate -> create organization
	OrgCreate = NewPermission("org:v1alpha1:create", false)

	// OrgDelete -> delete organization
	OrgDelete = NewPermission("org:v1alpha1:delete", true)

	// OrgUpdate -> update info of organization
	OrgUpdate = NewPermission("org:v1alpha1:update", true)

	// OrgInvite -> invite a member to organization
	OrgInvite = NewPermission("org:v1alpha1:invite", true)
)
