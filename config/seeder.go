package config

import "GoTransact/models"

func Seed() {
	// Seeder payment status
	var paymentM []models.PaymentStatus
	DB.Find(&paymentM)
	if len(paymentM) == 0 {
		param := []models.PaymentStatus{
			{PaymentStatus: "00", PaymentStatusDescription: "Transaksi Berhasil"},
			{PaymentStatus: "01", PaymentStatusDescription: "Transaksi Gagal"},
			{PaymentStatus: "02", PaymentStatusDescription: "Transaksi Pending"},
			{PaymentStatus: "03", PaymentStatusDescription: "Transaksi Ditolak"},
		}
		DB.Create(&param)
	}
}
