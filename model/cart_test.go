package model

import (
	"fmt"
	"testing"

	"github.com/relax-space/go-kit/test"
)

func Test_GetAllCarts(t *testing.T) {
	totalCount, carts, err := Cart{}.GetAllCarts(0, 1)
	fmt.Println(carts)
	test.Ok(t, err)
	test.Equals(t, int64(1), totalCount)
	test.Equals(t, 1, len(carts))
	test.Equals(t, float64(22), carts[0].ListPrice)
	test.Equals(t, float64(25), carts[0].SalePrice)
	test.Equals(t, float64(100), carts[0].Quantity)

	test.Equals(t, float64(11), carts[0].Items[0].ListPrice)
	test.Equals(t, float64(13), carts[0].Items[0].SalePrice)

}

func Test_GetCart(t *testing.T) {
	cart, err := Cart{}.GetCart(int64(1))
	test.Ok(t, err)
	test.Equals(t, float64(2), cart.ListPrice)
	test.Equals(t, float64(1), cart.SalePrice)
	test.Equals(t, float64(100), cart.Quantity)
}

func Test_ClearCart(t *testing.T) {
	cart, err := Cart{}.ClearCart(int64(1))
	test.Ok(t, err)
	test.Equals(t, float64(0), cart.ListPrice)
	test.Equals(t, float64(0), cart.SalePrice)
	test.Equals(t, float64(0), cart.Quantity)
}
