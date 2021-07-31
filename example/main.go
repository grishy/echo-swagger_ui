package main

import (
	_ "embed"
	"net/http"

	"github.com/labstack/echo/v4"

	swaggerui "github.com/grishy/echo-swaggerui"
)

//go:embed swagger.json
var spec []byte

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Group("/swagger", swaggerui.Handler(spec))

	e.Logger.Fatal(e.Start(":1323"))
}
