package main

import (
	"flag"
	"net/http"

	"github.com/labstack/echo"
)

var (
	httpAddr = flag.String("http.addr", ":5000", "HTTP listen address")
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error { return c.String(http.StatusOK, "pong") })

}
