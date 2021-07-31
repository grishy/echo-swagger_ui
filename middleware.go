package echoswaggerui

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

//go:embed swagger_ui
var uiFiles embed.FS

func Handler(specYAML []byte) echo.MiddlewareFunc {
	httpFS := http.FS(uiFiles)

	return middleware.StaticWithConfig(middleware.StaticConfig{
		Skipper:    middleware.DefaultSkipper,
		Root:       "swagger_ui",
		Index:      "index.html",
		HTML5:      false,
		Browse:     false,
		IgnoreBase: false,
		Filesystem: httpFS,
	})
}
