package model

type Sku struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	ListPrice float64 `json:"list_price"`
}

func (a *Sku) Get(skuId int64) (cart *Cart, err error) {
	return
}
func CreateSku(skus []Sku) (err error) {
	_, err = db.Insert(skus)
	return
}
