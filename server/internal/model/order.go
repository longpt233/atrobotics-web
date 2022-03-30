package model

//Order --> Model to entity Order
type Order struct {
	User      User    `gorm:"foreignkey:UserID"`
	Product   Product `gorm:"foreignkey:ProductID"`
	UserID    uint
	ProductID uint
	Quantity  int `json:"quantity"`
}

//TableName --> Table for Order Model
func (Order) TableName() string {
	return "orders"
}
