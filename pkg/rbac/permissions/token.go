package permissions

var (
	TokenIssue  = NewPermission("token:v1alpha1:issue", NamespaceUser, NamespaceOrg)
	TokenDelete = NewPermission("token:v1alpha1:delete", NamespaceUser, NamespaceOrg)

	TokenPermissions = []*Permission{TokenIssue, TokenDelete}
)
