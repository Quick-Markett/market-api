package models

type OrderItem struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	OrderId    int     `json:"order_id"`
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`

	Product Product `gorm:"foreignKey:ProductId;references:ID;references:ID"`
}

var OrderItems []OrderItem

func (OrderItem) TableName() string {
	return "tbOrderItems"
}
