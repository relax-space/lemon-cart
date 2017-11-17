package model

import "fmt"

var (
	SkuNotFoundError = fmt.Errorf("sku is not found.")
)

type Sku struct {
	Id        int64   `json:"id"`
	Name      string  `json:"name"`
	ListPrice float64 `json:"list_price"`
	SalePrice float64 `json:"sale_price"`
}

func (Sku) Get(skuId int64) (sku *Sku, err error) {
	sku = &Sku{}
	has, err := db.ID(skuId).Get(sku)
	if !has {
		err = SkuNotFoundError
		return
	}
	return
}
func (Sku) CreateSkus(skus []Sku) (err error) {
	_, err = db.Insert(&skus)
	return
}
