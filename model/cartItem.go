package model

type CartItem struct {
	Id        int64   `json:"id"`
	CartId    int64   `json:"cart_id"`
	Sku       Sku    `json:"sku,omitempty" xorm:"varchar(1024)"`
	Quantity  float64 `json:"quantity"`
	ListPrice float64 `json:"list_price"`
	SalePrice float64 `json:"sale_price"`
	Discount  float64 `json:"discount"`
}
