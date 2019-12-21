package permissions

var (
	TokenIssue  = NewPermission("token:v1alpha1:issue", true)
	TokenDelete = NewPermission("token:v1alpha1:delete", true)
)
