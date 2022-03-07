package controller

import (
	"net/http"

	"go-message-api/model"

	"github.com/labstack/echo/v4"
)

func CreateMemo(c echo.Context) error {
	memo := &model.Memo{
		Content: "ああああ",
	}

	return c.JSON(http.StatusOK, memo)
}
