package model

type basicService struct {
}

func (s *basicService) CreateCart() (cart *Cart, err error) {
	cart, err = Cart{}.create()
	return
}

func (s *basicService) GetAllCarts() (carts []Cart, err error) {
	carts, err = Cart{}.getAll()
	return
}

func (s *basicService) GetCart(cartId int64) (cart *Cart, err error) {
	cart, err = Cart{}.get(cartId)
	return
}
func (s *basicService) ClearCart(cartId int64) (cart *Cart, err error) {
	cart, err = Cart{}.clear(cartId)
	return
}

func (s *basicService) AddItems(cartId int64) (cart *Cart, err error) {
	cart, err = Cart{}.clear(cartId)
	return
}
