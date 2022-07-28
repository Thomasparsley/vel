package identity

import (
	"github.com/Thomasparsley/vel/database/fields"
)

const TableName_Users = "vel__users"

type User struct {
	fields.IdField[User]
	Username string `gorm:"size:64;index"`
	Email    string `gorm:"size:320;index"`
	Password string `gorm:"size:128"`
	Enabled  bool   `gorm:"default:true"`
	Admin    bool   `gorm:"default:false"`
	Role     string `gorm:"size:12"`
	fields.CreatedAtField
	fields.UpdatedAtField
}

func (User) TableName() string {
	return TableName_Users
}

func (user User) IsAdmin() bool {
	return user.Admin
}

func (user User) RoleName() RoleName {
	return RoleName(user.Role)
}

func (user User) HasRole(roleName RoleName) bool {
	return user.RoleName() == roleName
}

func (user User) HasPermission(permissionName PermissionName) bool {
	if user.IsAdmin() {
		return true
	}

	val, ok := GetPermissionsMap()[user.RoleName()][permissionName]
	if ok && bool(val) {
		return true
	}

	return false
}
