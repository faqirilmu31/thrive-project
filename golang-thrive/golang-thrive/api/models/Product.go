package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

type Product struct {
	ProductId   int32     `gorm:"primary_key;auto_increment" json:"product_id"`
	DisplayName string    `gorm:"size:255;not null;" json:"display_name"`
	Category    string    `gorm:"size:255;not null;" json:"category"`
	Price       float32   `gorm:"not null" json:"price"`
	Quantity    int32     `gorm:"not null;" json:"quantity"`
	Description string    `gorm:"size:255;not null;" json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (p *Product) FindAllProduct(db *gorm.DB) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(&Product{}).Limit(100).Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	return &products, nil
}

func (p *Product) FindProductByName(db *gorm.DB, keyword string) (*[]Product, error) {
	var err error
	products := []Product{}
	err = db.Debug().Model(Product{}).Where("display_name LIKE ?", "%"+keyword+"%").Find(&products).Error
	if err != nil {
		return &[]Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &[]Product{}, errors.New("Product Not Found")
	}
	return &products, nil
}

func (p *Product) FindProductByID(db *gorm.DB, keyword int32) (*Product, error) {
	var err error
	products := Product{}
	err = db.Debug().Model(Product{}).Where("product_id = ?", keyword).Find(&products).Error
	if err != nil {
		return &Product{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Product{}, errors.New("Product Not Found")
	}
	return &products, nil
}

func (p *Product) UpdateUnit(db *gorm.DB, product_id int32) error {
	var err error
	db = db.Debug().Model(&Product{}).Where("product_id = ?", product_id).Take(&Product{}).UpdateColumns(
		map[string]interface{}{
			"Quantity":   p.Quantity,
			"updated_at": time.Now(),
		},
	)
	err = db.Debug().Model(&Product{}).Where("product_id = ?", product_id).Take(&Product{}).Error
	if err != nil {
		return err
	}
	return nil
}
