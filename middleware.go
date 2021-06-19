package echoswaggerui

import (
	"embed"
	"io/fs"
	"net/http"

	"github.com/labstack/echo/v4"
)

//go:embed swagger_ui
var uiFiles embed.FS

func Handler(e *echo.Echo, specYAML []byte) error {
	strip, err := fs.Sub(uiFiles, "swagger_ui")
	if err != nil {
		return err
	}

	httpFS := http.FS(strip)
	assetHandler := http.FileServer(httpFS)
	_ = assetHandler

	e.GET("/swagger/swagger.yaml", func(ctx echo.Context) error {
		return ctx.Blob(http.StatusOK, "application/x-yaml", specYAML)
	})

	e.GET("/swagger/*", echo.WrapHandler(http.StripPrefix("/swagger/", assetHandler)))

	return nil
}
