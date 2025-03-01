package models

type Customer struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	CustomerPan  string `gorm:"type:varchar(20);unique;not null" json:"customer_pan"`
	CustomerName string `gorm:"type:varchar(100);not null" json:"customer_name"`
}
