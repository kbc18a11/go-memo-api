package main

import (
	"fmt"
	"go-message-api/config"
	"go-message-api/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	err := config.SetUp()

	if err != nil {
		fmt.Errorf("DBを接続できませんでした。")
		return
	}

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.POST("/memo", controller.CreateMemo)
	e.Logger.Fatal(e.Start(":1323"))
}
