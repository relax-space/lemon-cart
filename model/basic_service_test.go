package model

// type Service interface {
// 	CreateCart(ctx context.Context) (*Cart, error)
// 	GetAllCarts(ctx context.Context, skipCount, maxResultCount int) (*api.ArrayResult, error)
// 	GetCart(ctx context.Context, cartId int64) (*Cart, error)
// 	ClearCart(ctx context.Context, cartId int64) (*Cart, error)
// 	RemoveCart(ctx context.Context, cartId int64) error
// 	AddItems(ctx context.Context, cartId int64, items []CartItemEditRequest) (*Cart, error)
// 	RemoveItem(ctx context.Context, cartId int64, item CartItemEditRequest) (*Cart, error)
// 	EditCartItemOffer(ctx context.Context, cartId int64, item CartItemEditRequest) (*Cart, error)
// 	AddOptionalOffer(ctx context.Context, cartId, offerId int64) (*Cart, error)
// 	ClearOptionalOffer(ctx context.Context, cartId int64) (*Cart, error)
// 	SetInfo(ctx context.Context, cartId int64, info map[string]interface{}) (*Cart, error)
// 	SetCustomer(ctx context.Context, cartId int64, brandCode, customerNo, mobile string) (*Cart, error)
// 	SetCoupon(ctx context.Context, cartId int64, couponNo string) (*Cart, error)
// 	SetPayment(ctx context.Context, cartId int64, payment Payment) (*Cart, error)
// 	SetUseMileage(ctx context.Context, cartId int64, amount int) (*Cart, error)
// 	SetSalesman(ctx context.Context, cartId, salesmanId int64) (*Cart, error)
// 	SaveCart(ctx context.Context, cart *Cart) (*Cart, error)
// 	Calculate(ctx context.Context, cart *Cart) (*Cart, error)
// }
