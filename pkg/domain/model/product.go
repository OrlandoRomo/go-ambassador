package model

type Product struct {
	ID          uint    `json:"product_id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

type SearchProduct struct {
	Search string
	Sort   string
	Page   int
	Result []*Product
}
