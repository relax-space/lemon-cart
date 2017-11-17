package model

import (
	"github.com/go-xorm/xorm"
)

var (
	db *dbEngine
)

func Init(dbXorm *xorm.Engine) {
	db = &dbEngine{dbXorm}
	db.Sync(new(Cart), new(CartItem), new(Sku))
}

type dbEngine struct{ *xorm.Engine }
