package user

import (
	"time"

	"gorm.io/gorm"

	"github.com/Thomasparsley/vel/database"
)

const (
	TableName_Users = "users"
)

type User struct {
	ID        uint64     `gorm:"primaryKey;->;index"`
	Username  string     `gorm:"size:64;index"`
	Email     string     `gorm:"size:320;index"`
	Password  string     `gorm:"size:128"`
	Admin     bool       `gorm:"default:false"`
	Enabled   bool       `gorm:"default:true"`
	Role      string     `gorm:"size:3"`
	CreatedAt time.Time  `gorm:"autoCreateTime;not null;index"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime"`
}

func (u User) PK() uint64 {
	return u.ID
}

func (User) TableName() string {
	return TableName_Users
}

func (User) Object(db *gorm.DB) database.Object[uint64, User] {
	return database.NewObject[uint64, User](db)
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

func (u User) HasPermissions(name PermissionName, permissions Permissions) bool {
	if u.IsAdmin() {
		return true
	}

	val, ok := permissions[u.RoleName()][name]
	if ok && bool(val) {
		return true
	}

	return false
}
