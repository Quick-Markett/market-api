package models

type OrderItem struct {
	ID         uint    `json:"id"`
	OrderId    uint    `json:"order_id"`
	ProductId  uint    `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`

	Product Product `gorm:"foreignKey:ProductId;references:ID;references:ID"`
}

var OrderItems []OrderItem

func (OrderItem) TableName() string {
	return "tbOrderItems"
}
