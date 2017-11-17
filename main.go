package main

import (
	"flag"
	"fmt"
	"lemon-cart-api/controller"
	"lemon-cart-api/model"
	"net/http"
	"os"
	"runtime"

	"github.com/go-xorm/xorm"
	"github.com/labstack/echo"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/go-sql-driver/mysql"
)

var (
	httpAddr = flag.String("http.addr", ":5000", "HTTP listen address")
	connStr  = flag.String("CONN_STR", os.Getenv("CONN_STR"), "CONN_STR")
	dbDrive  = flag.String("DB_DRIVE", os.Getenv("DB_DRIVE"), "DB_DRIVE")
)

func init() {
	if *dbDrive == "sqlite3" {
		runtime.GOMAXPROCS(1)
	}
	db, err := xorm.NewEngine(*dbDrive, *connStr)
	if err != nil {
		panic(fmt.Errorf("Database open error: %s \n", err))
	}
	db.ShowSQL(true)
	fmt.Println(*dbDrive, *connStr)
	model.Init(db)
	model.Seed_sku("model/seed_sku.yaml")
	model.Seed_cart("model/seed_cart.yaml")
}

func main() {

	e := echo.New()
	e.GET("/ping", func(c echo.Context) error { return c.String(http.StatusOK, "pong") })
	v1 := e.Group("/v1")
	v1.GET("/carts", controller.GetAllCarts)
	v1.GET("/carts/:id", controller.GetCart)
	v1.DELETE("/carts/:id", controller.RemoveCart)
	v1.POST("/carts", controller.CreateCart)
	v1.POST("/carts/:id/items", controller.AddItems)
	v1.DELETE("/carts/:id/items", controller.RemoveItem)
	e.Start(":5000")
}
