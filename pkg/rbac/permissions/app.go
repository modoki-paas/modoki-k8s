package permissions

var (
	AppCreate = NewPermission("app:v1alpha1:create", true)
	AppUpdate = NewPermission("app:v1alpha1:update", true)
	AppDelete = NewPermission("app:v1alpha1:delete", true)
)
