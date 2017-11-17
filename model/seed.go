package model

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func seed_cart(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var carts []Cart
	if err := yaml.Unmarshal(data, &carts); err != nil {
		panic(err)
	}
	fmt.Println(carts)
	Cart{}.CreateCarts(carts)
}

func seed_sku(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var dtos []Sku
	if err := yaml.Unmarshal(data, &dtos); err != nil {
		panic(err)
	}
	Sku{}.CreateSkus(dtos)
}

func Seed_cart(path string) {
	Cart{}.CreateCart()
}

func Seed_sku(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var dtos []Sku
	if err := yaml.Unmarshal(data, &dtos); err != nil {
		panic(err)
	}
	Sku{}.CreateSkus(dtos)
}
