package model

type Product struct {
	ProductId	 int       `gorm:"primary_key;auto_increment" json:"product_id"`
	ProductName  string    `gorm:"size:255;not null;unique" json:"product_name"`
	Category 	 string    `gorm:"size:255;not null;" json:"category"`
	Price 		 int       `gorm:"not null;" json:"price"`
	Unit 		 int    `gorm:"size:255;not null;" json:"unit"`
	Description  string    `gorm:"size:255;not null;" json:"description"`
}