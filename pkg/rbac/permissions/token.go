package permissions

var (
	TokenIssue        = NewPermission("token:v1alpha1:issue", NamespaceUser, NamespaceOrg)
	TokenIssueUnbound = NewPermission("token:v1alpha1:unbound_issue", NamespaceUser, NamespaceOrg) // For OpenID Connect
	TokenBindTarget   = NewPermission("token:v1alpha1:bind_target", NamespaceUser, NamespaceOrg)   // For OpenID Connect
	TokenDelete       = NewPermission("token:v1alpha1:delete", NamespaceUser, NamespaceOrg)
	TokenValidate     = NewPermission("token:v1alpha:validate")

	TokenPermissions = []*Permission{TokenIssue, TokenDelete}
)
