package user

type PermissionValue bool
type PermissionName string
type RoleName string
type RolePermissions map[PermissionName]PermissionValue
type Permissions map[RoleName]RolePermissions
