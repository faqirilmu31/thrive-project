package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Cart struct {
	CartId    int32     `gorm:"primary_key;auto_increment" json:"cart_id"`
	ProductId int32     `gorm:"size:255;not null;" json:"product_id"`
	UserId    uint32    `gorm:"size:255;not null;" json:"user_id"`
	Total     int32     `gorm:"size:255;not null;" json:"total"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (c *Cart) Validate() error {

	if c.ProductId == 0 {
		return errors.New("required Product Id")
	}
	if c.Total == 0 {
		return errors.New("required Total")
	}
	return nil
}

func (c *Cart) AddCart(db *gorm.DB) (*Cart, error) {
	//var err error
	err := db.Debug().Create(&c).Error
	if err != nil {
		return &Cart{}, err
	}
	return c, nil
}

func (c *Cart) FindCartByID(db *gorm.DB, keyword int32) (*Cart, error) {
	var err error
	cart := Cart{}
	err = db.Debug().Model(Cart{}).Where("cart_id = ?", keyword).Find(&cart).Error
	if err != nil {
		return &Cart{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Cart{}, errors.New("Cart Not Found")
	}
	return &cart, nil
}

func (c *Cart) DeleteCart(db *gorm.DB, cart_id int32) error {

	db = db.Debug().Model(&Cart{}).Where(" cart_id = ?", cart_id).Take(&Cart{}).Delete(&Cart{})

	if db.Error != nil {
		if gorm.IsRecordNotFoundError(db.Error) {
			return errors.New("Cart not found")
		}
		return db.Error
	}
	return nil
}
