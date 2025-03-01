package models

type PaymentStatus struct {
	ID                       uint   `gorm:"primaryKey" json:"id"`
	PaymentStatus            string `gorm:"type:varchar(10)not null" json:"payment_status"`
	PaymentStatusDescription string `gorm:"type:varchar(100);not null" json:"payment_status_description"`
}
