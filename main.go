package main

import (
	"flag"
	"fmt"
	"go-message-api/config"
	"go-message-api/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	isBatch := flag.Bool("isBatch", false, "バッチの実行について")
	batchType := flag.String("batchType", "", "実行するバッチの種類")
	flag.Parse()

	if *isBatch {
		// バッチ処理である場合
		doBatch(*batchType)
		return
	}

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

/*
バッチの実行
*/
func doBatch(batchType string) {
	switch batchType {
	case "migration":
		fmt.Println(batchType)
		break

	default:
		fmt.Errorf("存在しないバッチです。")
	}
}
