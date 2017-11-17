package model

import (
	"fmt"
	"time"

	"github.com/go-xorm/xorm"
	"github.com/relax-space/go-kit/base"
)

var (
	CartNotFoundError = fmt.Errorf("cart is not found.")
)

type Cart struct {
	Id        int64      `json:"id"`
	Items     []CartItem `json:"items" xorm:"-"`
	ListPrice float64    `json:"listPrice"`
	SalePrice float64    `json:"salePrice"`
	Quantity  float64    `json:"quantity"`

	CreatedAt time.Time `json:"createdAt" xorm:"created"`
	UpdateAt  time.Time `json:"updatedAt" xorm:"updated"`
}

func (Cart) CreateCarts(carts []Cart) (err error) {
	for _, v := range carts {
		err = v.calculateAndSave()
		if err != nil {
			return
		}
	}
	return
}

func (Cart) CreateCart() (cart *Cart, err error) {
	cart = &Cart{}
	_, err = db.Insert(cart)
	return
}

func (Cart) GetAllCarts(skipCount, maxResultCount int) (totalCount int64, carts []Cart, err error) {
	queryBuilder := func(session *xorm.Session) *xorm.Session {
		return session.Where("cart.sale_price != 0")
	}

	totalCount, err = queryBuilder(db.NewSession()).Count(&Cart{})
	if err != nil {
		return
	}

	var cartItemExtends []struct {
		Cart     `xorm:"extends"`
		CartItem `xorm:"extends"`
	}

	if err := queryBuilder(db.Table("cart").Select("cart.*, cart_item.*")).
		Join("INNER", "cart_item", "cart.id = cart_item.cart_id").
		Desc("cart.id").
		Limit(maxResultCount, skipCount).
		Find(&cartItemExtends); err != nil {
		return 0, nil, err
	}
	getExistCart := func(id int64) *Cart {
		for i := range carts {
			if carts[i].Id == id {
				return &carts[i]
			}
		}
		return nil
	}
	for _, t := range cartItemExtends {
		c := getExistCart(t.Cart.Id)
		if c == nil {
			carts = append(carts, t.Cart)
			c = &carts[len(carts)-1]
		}
		c.Items = append(c.Items, t.CartItem)
	}

	return
}

func (Cart) GetCart(id int64) (cart *Cart, err error) {
	cart, err = Cart{}.get(id)
	return
}
func (Cart) ClearCart(cartId int64) (cart *Cart, err error) {
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
func (Cart) RemoveCart(cartId int64) (err error) {
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

func (Cart) AddItems(cartId int64, items []CartItemEditRequest) (cart *Cart, err error) {
	cart, err = Cart{}.get(cartId)
	if err != nil {
		return
	}
	for _, item := range items {
		var skuInfo *Sku
		skuInfo, err = Sku{}.Get(item.SkuId)
		if err != nil {
			return
		}
		fmt.Println(skuInfo)
		err = cart.addItem(*skuInfo, item.Quantity)
		if err != nil {
			return
		}
	}
	err = cart.calculateAndSave()
	return
}
func (Cart) RemoveItem(cartId int64, item CartItemEditRequest) (cart *Cart, err error) {
	cart, err = Cart{}.get(cartId)
	if err != nil {
		return
	}
	fmt.Printf("1111%+v", cart)
	err = cart.removeItem(item.SkuId, item.Quantity)
	if err != nil {
		return
	}
	fmt.Printf("222%+v", cart)
	err = cart.calculateAndSave()
	if err != nil {
		return
	}
	return
}

func (c *Cart) SaveCart() (err error) {
	err = c.calculateAndSave()
	return
}

func (c *Cart) removeItem(skuId int64, quantity float64) (err error) {
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
			c.Items = []CartItem{}
		case len(c.Items)-1 == emptyItemIndex:
			c.Items = c.Items[:emptyItemIndex]
		default:
			c.Items = append(c.Items[:emptyItemIndex], c.Items[emptyItemIndex+1:]...)
		}
	}
	return
}

func (Cart) get(id int64) (cart *Cart, err error) {
	cart = &Cart{}
	has, err := db.Id(id).Get(cart)
	if !has {
		err = CartNotFoundError
		return
	}

	err = db.Where("cart_id = ?", cart.Id).Find(&cart.Items)
	if err != nil {
		return
	}
	return
}

func (c *Cart) calculateAndSave() (err error) {
	err = c.calculate(&base.PriceSetting{
		RoundDigit:    2,
		RoundStrategy: "round",
		Currency:      "CYN",
	})
	fmt.Printf("3333%+v", c)

	if err != nil {
		return
	}
	err = c.save()
	if err != nil {
		return
	}
	return
}

func (cart *Cart) save() (err error) {
	cols := []string{"list_price", "sale_price", "quantity"}
	if cart.Id == 0 {
		if _, err = db.InsertOne(cart); err != nil {
			return
		}
	} else {
		if _, err = db.ID(cart.Id).Cols(cols...).Update(cart); err != nil {
			return
		}
	}
	fmt.Printf("444%+v", cart)
	var cartItemIds []interface{}
	for i := range cart.Items {
		if cart.Items[i].Id == 0 {
			cart.Items[i].CartId = cart.Id
			if _, err = db.InsertOne(&cart.Items[i]); err != nil {
				return
			}
		} else {
			if _, err = db.ID(cart.Items[i].Id).Cols("sku", "quantity", "list_price", "sale_price", "discount").Update(&cart.Items[i]); err != nil {
				return
			}
		}
		cartItemIds = append(cartItemIds, cart.Items[i].Id)
	}
	fmt.Printf("5555%+v", cart)
	db.Where("cart_id = ?", cart.Id).NotIn("id", cartItemIds...).Delete(&CartItem{})
	fmt.Printf("6666%+v", cart)
	return
}

func (c *Cart) addItem(sku Sku, quantity float64) (err error) {
	for i, t := range c.Items {
		if t.Sku.Id == sku.Id {
			c.Items[i].Quantity += quantity
			return
		}
	}
	c.Items = append(c.Items, CartItem{
		Sku:      sku,
		Quantity: quantity,
	})
	return
}

func (c *Cart) calculate(priceSettings *base.PriceSetting) (err error) {
	c.ListPrice = 0
	c.SalePrice = 0
	c.Quantity = 0
	for i, t := range c.Items {
		c.Items[i].ListPrice = base.ToFixed(base.ToFixed(t.Sku.ListPrice, priceSettings)*t.Quantity, priceSettings)
		c.Items[i].SalePrice = base.ToFixed(base.ToFixed(t.Sku.SalePrice, priceSettings)*t.Quantity, priceSettings)
		c.ListPrice += c.Items[i].ListPrice
		c.SalePrice += c.Items[i].SalePrice
		c.Quantity += t.Quantity
	}
	return
}
