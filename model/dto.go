package model

type CartItemEditRequest struct {
	SkuId    int64   `json:"sku_id"`
	Quantity float64 `json:"quantity"`
	ItemId   int64   `json:"item_id"`
}
