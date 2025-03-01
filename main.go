package main

import (
	"GoTransact/config"
	"GoTransact/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	// connect database
	config.InitDB()
	// Migrate
	config.DB.AutoMigrate(&models.Customer{})
	config.DB.AutoMigrate(&models.Merchant{})
	config.DB.AutoMigrate(&models.PaymentStatus{})
	config.DB.AutoMigrate(&models.Transaction{})

	// initialize echo
	e := echo.New()

	// Seeder payment status
	var paymentM []models.PaymentStatus
	config.DB.Find(&paymentM)
	if len(paymentM) == 0 {
		param := []models.PaymentStatus{
			{PaymentStatus: "00", PaymentStatusDescription: "Transaksi Berhasil"},
			{PaymentStatus: "01", PaymentStatusDescription: "Transaksi Gagal"},
			{PaymentStatus: "02", PaymentStatusDescription: "Transaksi Pending"},
			{PaymentStatus: "03", PaymentStatusDescription: "Transaksi Ditolak"},
		}
		config.DB.Create(&param)
	}

	e.GET("/hello", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
