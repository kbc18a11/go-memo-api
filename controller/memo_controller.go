package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Memo struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
}

func CreateMemo(c echo.Context) error {
	memo := &Memo{
		Id:      1,
		Content: "ああああ",
	}

	return c.JSON(http.StatusOK, memo)
}
