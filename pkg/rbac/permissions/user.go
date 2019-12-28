package permissions

var (
	UserCreate = NewPermission("user:v1alpha1:create")
	UserList   = NewPermission("user:v1alpha1:list")
	UserDelete = NewPermission("user:v1alpha1:delete", NamespaceUser)
	UserGet    = NewPermission("user:v1alpha1:get", NamespaceUser)
	UserUpdate = NewPermission("user:v1alpha1:update", NamespaceUser)

	UserPermissions = []*Permission{UserCreate, UserGet, UserDelete, UserUpdate, UserList}
)
