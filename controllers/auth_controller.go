package controllers

import (
	"GoTransact/middleware"
	"GoTransact/utils"
	"net/http"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

type UserValidate struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

func Login(c echo.Context) error {
	// request struct validation
	var user UserValidate

	// request params, and check body
	if err := c.Bind(&user); err != nil {
		utils.Logger.Error("Invalid request body")
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "invalid request body",
		})
	}

	// validation struc
	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		errors := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			errors[err.Field()] = "This field is" + " " + err.Tag() + " " + err.Param()
		}
		utils.Logger.Error(errors)
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": errors,
		})
	}

	// check email and password
	email := "admin@mail.com"
	password := "Admin123#"

	if user.Email != email || user.Password != password {
		utils.Logger.Warn("Invalid email or password")
		return c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "Invalid email or password",
		})
	}

	// Get generate token
	token, err := middleware.GenerateToken(user.Email)
	if err != nil {
		utils.Logger.Error("Failed to generate token")
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    http.StatusInternalServerError,
			"message": "Failed to generate token",
		})
	}

	// return success
	utils.Logger.Info("Login successfully")
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":      200,
		"message":   "Login successfully",
		"token":     token,
		"token_exp": time.Now().Add(time.Hour * 1),
	})
}
