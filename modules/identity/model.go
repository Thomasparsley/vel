package identity

import (
	"github.com/Thomasparsley/vel/database/fields"
)

const (
	TableName_Users = "velusers"
)

type User struct {
	fields.IdField[User]
	Username string `gorm:"size:64;index"`
	Email    string `gorm:"size:320;index"`
	Password string `gorm:"size:128"`
	Admin    bool   `gorm:"default:false"`
	Enabled  bool   `gorm:"default:true"`
	Role     string `gorm:"size:3"`
	fields.CreatedAtField
	fields.UpdatedAtField
}

func (User) TableName() string {
	return TableName_Users
}

func (u User) IsAdmin() bool {
	return u.Admin
}

func (u User) RoleName() RoleName {
	return RoleName(u.Role)
}

func (u User) HasRole(name RoleName) bool {
	return u.RoleName() == name
}

func (u User) HasPermission(name PermissionName, permissions Permissions) bool {
	if u.IsAdmin() {
		return true
	}

	val, ok := permissions[u.RoleName()][name]
	if ok && bool(val) {
		return true
	}

	return false
}
