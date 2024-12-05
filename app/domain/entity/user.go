package entity

import (
	"time"

	"github.com/google/uuid"
)

type UserRole int

const (
	UserRoleNormal UserRole = iota + 1
	UserRoleManager
	UserRoleAdministrator
	UserRoleSuperAdmin
)

const (
	userRole       = "user"
	managerRole    = "manager"
	adminRole      = "admin"
	superadminRole = "superadmin"
)

func GetUserRolePriority(role UserRole) uint32 {
	switch role {
	case UserRoleSuperAdmin:
		return 1
	case UserRoleAdministrator:
		return 2
	case UserRoleManager:
		return 3
	case UserRoleNormal:
		return 4
	default:
		return 5
	}
}
func (r UserRole) String() string {
	switch r {
	case UserRoleNormal:
		return userRole
	case UserRoleManager:
		return managerRole
	case UserRoleAdministrator:
		return adminRole
	case UserRoleSuperAdmin:
		return superadminRole
	default:
		return ""
	}
}

// GetUserRole returns an  user role as an enum type
func GetUserRole(role string) UserRole {
	switch role {
	case userRole:
		return UserRoleNormal
	case managerRole:
		return UserRoleManager
	case adminRole:
		return UserRoleAdministrator
	case superadminRole:
		return UserRoleSuperAdmin
	default:
		return 0
	}
}

// User is the struct for user entity
type User struct {
	UserID      uuid.UUID `gorm:"column:user_id"`
	UserName    string    `gorm:"column:username"`
	Password    string    `gorm:"column:password"`
	Email       string    `gorm:"column:email"`
	PhoneNumber string    `gorm:"column:phone_number"`
	AvatarLink  string    `gorm:"column:avatar_link"`
	RoleCode    UserRole  `gorm:"column:role_code"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
	IsDeleted   bool      `gorm:"column:is_deleted"`
}

// IsDeletedUser a user is deleted
func (u *User) IsDeletedUser() bool {
	return u.IsDeleted
}

// SuperAdmin to know if the user is a super admin
func (u *User) SuperAdmin() bool {
	return u.RoleCode == UserRoleSuperAdmin
}

// Admin to know if the user is an admin
func (u *User) Admin() bool {
	return u.RoleCode == UserRoleAdministrator
}

// Manager to know if the user is a manager
func (u *User) Manager() bool {
	return u.RoleCode == UserRoleManager
}

// User to know if the user is a user
func (u *User) User() bool {
	return u.RoleCode == UserRoleNormal
}

// ToMap returns a map of updatable fields
func (user User) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"role":        user.RoleCode,
		"avatar_link": user.AvatarLink,
	}
}
