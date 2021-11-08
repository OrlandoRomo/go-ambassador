package model

type Link struct {
	ID       uint      `json:"id"`
	Code     string    `json:"code"`
	UserID   uint      `json:"user_id"`
	User     User      `json:"user" gorm:"foreignKey:UserID"`
	Products []Product `json:"products" gorm:"many2many:link_products"`
	Orders   []Order   `json:"orders" gorm:"-"`
}

type LinkRequest struct {
	Products []int `json:"products"`
}
