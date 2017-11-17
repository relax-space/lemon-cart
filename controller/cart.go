package controller

import (
	"lemon-cart-api/model"
	"net/http"
	"strconv"

	rmode "github.com/relax-space/go-kit/model"

	"github.com/labstack/echo"
)

func GetAllCarts(c echo.Context) error {

	skipCount := c.QueryParam("skipCount")
	maxResultCount := c.QueryParam("maxResultCount")
	var skip, limit int64
	var err error
	if len(maxResultCount) == 0 {
		limit = 30
	} else {
		limit, err = strconv.ParseInt(maxResultCount, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
		}
	}
	if len(skipCount) == 0 {
		skip = 0
	} else {
		skip, err = strconv.ParseInt(skipCount, 10, 64)
		if err != nil {
			return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
		}
	}

	total, carts, err := (model.Cart{}).GetAllCarts(int(skip), int(limit))
	if err != nil {
		return c.JSON(http.StatusOK, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	result := rmode.ArrayResult{
		TotalCount: int(total),
		Items:      carts,
	}
	return c.JSON(http.StatusOK, rmode.Result{Success: true, Result: result})

}

func GetCart(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	card, err := model.Cart{}.GetCart(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, rmode.Result{Success: true, Result: card})

}

func RemoveCart(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	err = model.Cart{}.RemoveCart(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	return c.JSON(http.StatusNoContent, rmode.Result{Success: true})
}

func AddItems(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	var cartItem []model.CartItemEditRequest
	err = c.Bind(&cartItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}

	cart, err := model.Cart{}.AddItems(id, cartItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, rmode.Result{Success: true, Result: cart})
}

func RemoveItem(c echo.Context) error {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	var cartItem model.CartItemEditRequest
	err = c.Bind(&cartItem)
	if err != nil {
		return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}

	cart, err := model.Cart{}.RemoveItem(id, cartItem)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	return c.JSON(http.StatusOK, rmode.Result{Success: true, Result: cart})
}

func CreateCart(c echo.Context) error {
	// cart := new(model.Cart)
	// err := c.Bind(cart)
	// if err != nil {
	// 	return c.JSON(http.StatusBadRequest, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	// }
	cart, err := model.Cart{}.CreateCart()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, rmode.Result{Success: false, Error: rmode.Error{Message: err.Error()}})
	}
	return c.JSON(http.StatusCreated, rmode.Result{Success: true, Result: cart})

}
