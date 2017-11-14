package model

import "io/ioutil"
import "gopkg.in/yaml.v2"

// func seed_cart(path string) {
// 	data, err := ioutil.ReadFile(path)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var carts []Cart
// 	if err := yaml.Unmarshal(data, &carts); err != nil {
// 		panic(err)
// 	}
// 	CreateCart(carts)
// }

func seed_sku(path string) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var dtos []Sku
	if err := yaml.Unmarshal(data, &dtos); err != nil {
		panic(err)
	}
	CreateSku(dtos)
}
