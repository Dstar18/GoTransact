package controllers

import (
	"GoTransact/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetCustomers(c echo.Context) error {
	utils.Logger.Info("JWT Successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    200,
		"message": "JWT Successfully",
	})
}
