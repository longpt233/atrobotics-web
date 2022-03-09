package model

type Role struct{
	RoleID int `json:"role_id"  gorm:"primary_key"`
	RoleName string `json:"role_name" `
}

func (Role) TableName() string {
	return "roles"
}