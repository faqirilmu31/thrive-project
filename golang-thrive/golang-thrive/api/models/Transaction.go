package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Transaction struct {
	TransactionId string    `gorm:"size:255;not null;" json:"transaction_id"`
	ProductId     int32     `gorm:"size:255;not null;" json:"product_id"`
	UserId        uint32    `gorm:"size:255;not null;" json:"user_id"`
	CartId        int32     `gorm:"size:255;not null;" json:"cart_id"`
	Price         float32   `gorm:"size:255;not null;" json:"price"`
	Total         int32     `gorm:"size:255;not null;" json:"total"`
	PaymentStatus string    `gorm:"size:255;not null;" json:"payment_status"`
	CreatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt     time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func InsertTransaction(db *gorm.DB, transaction Transaction) (err error) {
	args := []interface{}{transaction.TransactionId, transaction.ProductId, transaction.UserId, transaction.CartId, transaction.Price, transaction.Total, transaction.PaymentStatus}
	err = db.Exec("INSERT INTO transactions(transaction_id, product_id, user_id, cart_id, price, total, payment_status) VALUES (?, ?, ?, ?, ?, ?, ?)", args...).Error
	return
}

func (t *Transaction) FindTransactionByID(db *gorm.DB, transaction_id string) (*[]Transaction, error) {
	var err error
	transaction := []Transaction{}
	err = db.Debug().Model(Product{}).Where("transaction_id = ?", transaction_id).Find(&transaction).Error
	if err != nil {
		return &[]Transaction{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Transaction{}, errors.New("Transaction Not Found")
	}
	return &transaction, nil
}
