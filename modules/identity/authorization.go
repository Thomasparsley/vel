package identity

type PermissionValue bool
type PermissionName string
type RoleName string
type RolePermissions map[PermissionName]PermissionValue
type Permissions map[RoleName]RolePermissions

var permissionsMap Permissions = nil

func SetPermissionsMap(permissions Permissions) {
	permissionsMap = permissions
}

func GetPermissionsMap() Permissions {
	return permissionsMap
}
