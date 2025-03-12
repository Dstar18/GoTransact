package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Transaction struct {
	ID              uint            `gorm:"primaryKey" json:"id"`
	RequestID       string          `gorm:"type:varchar(50);unique;not null" json:"request_id"`
	CustomerFKID    uint            `gorm:"not null" json:"customerfk_id"`     //FK
	MerchantFKID    uint            `gorm:"not null" json:"merchantfk_id"`     //FK
	PaymentStatusID uint            `gorm:"not null" json:"payment_status_id"` //FK
	Amount          decimal.Decimal `gorm:"type:decimal(15,2);not null"`
	TransactionAt   string          `gorm:"type:timestamptz;default:null" json:"transaction_at"`
	RRN             string          `gorm:"type:varchar(50);not null" json:"rrn"`
	BillNumber      string          `gorm:"type:varchar(50);not null" json:"bill_number"`
	CurrencyCode    string          `gorm:"type:varchar(50);not null" json:"currency_code"`
	CreatedAt       time.Time       `gorm:"type:timestamp;default:null" json:"created_at"`
}
