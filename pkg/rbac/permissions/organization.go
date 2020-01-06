package permissions

var (
	// OrgCreate -> create organization
	OrgCreate = NewPermission("org:v1alpha1:create")

	// OrgList -> list joining organizations
	OrgList = NewPermission("org:v1alpha1:list")

	// OrgListAll -> list all organizations
	OrgListAll = NewPermission("org:v1alpha1:list_all")

	// OrgGet -> get organization
	OrgGet = NewPermission("org:v1alpha1:get", NamespaceOrg)

	// OrgDelete -> delete organization
	OrgDelete = NewPermission("org:v1alpha1:delete", NamespaceOrg)

	// OrgUpdate -> update info of organization
	OrgUpdate = NewPermission("org:v1alpha1:update", NamespaceOrg)

	// OrgMemberManagement -> manage members in a organization
	OrgMemberManagement = NewPermission("org:v1alpha1:member_management", NamespaceOrg)

	// OrgMemberList -> list members in a organization
	OrgMemberList = NewPermission("org:v1alpha1:member_list", NamespaceOrg)

	// UserOrgGetRoleBinding -> get role bindings for user/org
	UserOrgGetRoleBinding = NewPermission("userorg:v1alpha1:get_role_binding")

	OrgPermissions = []*Permission{OrgCreate, OrgList, OrgListAll, OrgDelete, OrgUpdate, OrgMemberManagement, OrgMemberList}
)
