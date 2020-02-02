package permissions

var (
	TokenIssue    = NewPermission("token:v1alpha1:issue", NamespaceUser, NamespaceOrg)
	TokenDelete   = NewPermission("token:v1alpha1:delete", NamespaceUser, NamespaceOrg)
	TokenValidate = NewPermission("token:v1alpha:validate")

	TokenPermissions = []*Permission{TokenIssue, TokenDelete}
)
