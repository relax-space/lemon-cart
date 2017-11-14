package model

import (
	"flag"
	"fmt"
	"os"
	"runtime"

	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
)

var (
	connStr = flag.String("CONN_STR", os.Getenv("CONN_STR"), "CONN_STR")
	dbDrive = flag.String("DB_DRIVE", os.Getenv("DB_DRIVE"), "DB_DRIVE")
	db      *xorm.Engine
)

func init() {
	if *dbDrive == "sqlite3" {
		runtime.GOMAXPROCS(1)
	}
	var err error
	db, err = xorm.NewEngine(*dbDrive, *connStr)
	if err != nil {
		panic(fmt.Errorf("Database open error: %s \n", err))
	}
}
func Init(dbXorm *xorm.Engine) {
	db = dbXorm
	db.Sync(new(Cart), new(CartItem), new(Sku))
}
