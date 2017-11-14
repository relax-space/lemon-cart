package model

import (
	"fmt"
	"time"
)

var (
	CartNotFoundError = fmt.Errorf("cart is not found.")
)

type Cart struct {
	Id        int64      `json:"id"`
	Items     []CartItem `json:"items" xorm:"-"`
	ListPrice float64    `json:"list_price"`
	SalePrice float64    `json:"sale_price"`
	Quantity  float64    `json:"quantity"`

	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"updated_at"`
}

func (Cart) getAll() (carts []Cart, err error) {
	err = db.Find(&carts)
	return
}

func (Cart) get(id int64) (card *Cart, err error) {
	has, err := db.Id(id).Get(card)
	if !has {
		err = CartNotFoundError
	}
	return
}

func (Cart) create() (result *Cart, err error) {
	cart := Cart{}
	_, err = db.Insert(&cart)
	result = &cart
	return
}

func (Cart) clear(cartId int64) (cart *Cart, err error) {
	cart, err = Cart{}.get(cartId)
	if err != nil {
		return
	}
	_, err = db.Where("cart_id=?", cartId).Delete(&CartItem{})
	if err != nil {
		return
	}
	cart.ListPrice = 0
	cart.SalePrice = 0
	cart.Quantity = 0
	_, err = db.Id(cart.Id).Cols("list_price", "sale_price", "quantity").Update(cart)
	return
}

func (Cart) remove(cartId int64) (err error) {
	if _, err = (Cart{}).get(cartId); err != nil {
		return
	}
	if _, err = db.Delete(&CartItem{CartId: cartId}); err != nil {
		return
	}
	if _, err = db.ID(cartId).Delete(&Cart{}); err != nil {
		return
	}
	return
}
func (cart *Cart) save() (err error) {
	cols := []string{"list_price", "sale_price", "quantity"}
	if cart.Id == 0 {
		if _, err = db.InsertOne(cart); err != nil {
			return
		} else {
			if _, err = db.ID(cart.Id).Cols(cols...).Update(cart); err != nil {
				return
			}
		}
	}
	var cartItemIds []interface{}
	for i := range cart.Items {
		if cart.Items[i].Id == 0 {
			cart.Items[i].CartId = cart.Id
			if _, err = db.InsertOne(cart.Items[i]); err != nil {
				return
			}
		} else {
			if _, err = db.Cols("sku", "quantity", "list_price", "sale_price", "discount").Update(cart); err != nil {
				return
			}
		}
		cartItemIds = append(cartItemIds, cart.Items[i].Id)
	}
	db.Where("cart_id = ?", cart.Id).NotIn("id", cartItemIds...).Delete(&CartItem{})
	return
}

func (c *Cart) AddItem(sku Sku, quantity float64) (err error) {
	for i, t := range c.Items {
		if t.Sku.Id == sku.Id {
			c.Items[i].Quantity += quantity
			return
		}
	}
	return
}

func (c *Cart) RemoveItem(skuId int64, quantity float64) (err error) {
	emptyItemIndex := -1
	for i := range c.Items {
		if c.Items[i].Sku.Id == skuId {
			c.Items[i].Quantity -= quantity
			if c.Items[i].Quantity == 0 {
				emptyItemIndex = i
			}
		}
	}
	//Make delete Item top
	if emptyItemIndex != -1 {
		switch {
		case len(c.Items) == 1 && emptyItemIndex == 0:
		case len(c.Items)-1 == emptyItemIndex:
		default:
			c.Items = append(c.Items[:emptyItemIndex], c.Items[emptyItemIndex+1:]...)
		}
	}
	return
}
