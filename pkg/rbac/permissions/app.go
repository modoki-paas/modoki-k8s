package permissions

var (
	AppCreate = NewPermission("app:v1alpha1:create", NamespaceUser, NamespaceOrg)
	AppUpdate = NewPermission("app:v1alpha1:update", NamespaceUser, NamespaceOrg)
	AppDelete = NewPermission("app:v1alpha1:delete", NamespaceUser, NamespaceOrg)

	AppPermissions = []*Permission{AppCreate, AppUpdate, AppDelete}
)
