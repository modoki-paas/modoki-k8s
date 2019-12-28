package permissions

var (
	// OrgCreate -> create organization
	OrgCreate = NewPermission("org:v1alpha1:create")

	// OrgDelete -> delete organization
	OrgDelete = NewPermission("org:v1alpha1:delete", NamespaceOrg)

	// OrgUpdate -> update info of organization
	OrgUpdate = NewPermission("org:v1alpha1:update", NamespaceOrg)

	// OrgMemberManagement -> manage members in a organization
	OrgMemberManagement = NewPermission("org:v1alpha1:member_management", NamespaceOrg)

	OrgPermissions = []*Permission{OrgCreate, OrgDelete, OrgUpdate, OrgMemberManagement}
)
