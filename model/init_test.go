package model

import (
	"fmt"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

func init() {
	//os.Remove("cart.db")
	_db, err := xorm.NewEngine("sqlite3", ":memory:")
	_db.ShowSQL(true)
	if err != nil {
		panic(fmt.Errorf("Database open error: %s \n", err))
	}
	db = &dbEngine{_db}
	Init(_db)
	seed_sku("seed_sku.yaml")
	seed_cart("seed_cart.yaml")
}
