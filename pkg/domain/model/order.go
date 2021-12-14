package model

type Order struct {
	ID              uint        `json:"order_id"`
	TransactionID   string      `json:"transaction_id" gorm:"null"`
	UserID          uint        `json:"user_id"`
	Code            string      `json:"code"`
	AmbassadorEmail string      `json:"ambassador_email"`
	FirstName       string      `json:"first_name"`
	LastName        string      `json:"last_name"`
	Email           string      `json:"email"`
	Address         string      `json:"address" gorm:"null"`
	City            string      `json:"city" gorm:"null"`
	Country         string      `json:"country" gorm:"null"`
	Zip             string      `json:"zip" gorm:"null"`
	IsCompleted     bool        `json:"is_completed" gorm:"default:false"`
	Total           float64     `json:"total" gorm:"-"`
	OrderItems      []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	ID                uint    `json:"order_item_id"`
	OrderID           uint    `json:"order_id"`
	ProductTitle      string  `json:"product_title"`
	Price             float64 `json:"price"`
	Quantity          uint    `json:"quantity"`
	AdminRevenue      float64 `json:"admin_revenue"`
	AmbassadorRevenue float64 `json:"ambassador_revenue"`
}

func (o *Order) GetTotal() float64 {
	total := 0.0
	for _, orderItem := range o.OrderItems {
		total += orderItem.Price * float64(orderItem.Quantity)
	}
	return total
}

type OrderRequest struct {
	Code      string           `json:"code"`
	FirstName string           `json:"first_name"`
	LastName  string           `json:"last_name"`
	Email     string           `json:"email"`
	Address   string           `json:"address"`
	Country   string           `json:"country"`
	City      string           `json:"city"`
	Zip       string           `json:"zip"`
	Products  []map[string]int `json:"products"`
}
