package permissions

var (
	UserGet           = NewPermission("user:v1alpha1:get", false)
	UserProfileGet    = NewPermission("user:v1alpha1:profile_get", false)
	UserProfileUpdate = NewPermission("user:v1alpha1:profile_update", false)
	UserDelete        = NewPermission("user:v1alpha1:delete", false)
)
