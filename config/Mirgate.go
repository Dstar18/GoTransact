package config

import "GoTransact/models"

func Migrate() {
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Merchant{})
	DB.AutoMigrate(&models.PaymentStatus{})
	DB.AutoMigrate(&models.Transaction{})
}
