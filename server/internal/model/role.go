package model

type Role struct {
	RoleID   string `json:"id"  gorm:"primaryKey" gorm:"column:role_id"`
	RoleName string `json:"name" gorm:"column:role_name"`
}

func (Role) TableName() string {
	return "roles"
}
