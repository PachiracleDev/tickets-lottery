package main

import (
	"app/config"
	"app/shared"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	config.LoadEnv()
	shared.ConnectDB()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":6000"))
}
