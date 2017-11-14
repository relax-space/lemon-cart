package model

import (
	"fmt"
	"os"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	dbName := "cart.db"
	os.Remove(dbName)
	dbXorm, err := xorm.NewEngine("sqlite3", dbName)
	if err != nil {
		panic(fmt.Errorf("Database open error:%s \n", err))
	}
	Init(dbXorm)
	db = dbXorm
	seed_sku("seed_sku.yaml")
	//seed_cart("seed_cart.yaml")
}
