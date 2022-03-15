package model

type Role struct {
	RoleID   string    `json:"role_id"  gorm:"primaryKey"`
	RoleName string `json:"role_name" `
}

func (Role) TableName() string {
	return "roles"
}
