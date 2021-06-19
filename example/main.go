package main

import (
	_ "embed"
	"log"
	"net/http"

	swaggerui "github.com/grishy/echo-swaggerui"
	"github.com/labstack/echo/v4"
)

//go:embed swagger.json
var spec []byte

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	err := swaggerui.Handler(e, spec)
	if err != nil {
		log.Panic(err)
	}

	e.Logger.Fatal(e.Start(":1323"))
}
