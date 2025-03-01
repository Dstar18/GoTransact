package models

type Merchant struct {
	ID           uint   `gorm:"primaryKey" json:"id"`
	MerchantID   string `gorm:"type:varchar(50);unique;not null" json:"merchant_id"`
	MerchantName string `gorm:"type:varchar(100);not null" json:"merchant_name"`
	MerchantCity string `gorm:"type:varchar(100);not null" json:"merchant_city"`
}
