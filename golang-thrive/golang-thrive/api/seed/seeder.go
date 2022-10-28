package seed

import (
	"golang-thrive/api/models"
	"log"

	// "gorm.io/gorm"
	"github.com/jinzhu/gorm"
)

var users = []models.User{
	models.User{
		Username: "admin",
		Email:    "admin@gmail.com",
		Password: "password",
		UserType: "admin",
	},
	models.User{
		Username: "mayank134",
		Email:    "mayang@gmail.com",
		Password: "password",
		UserType: "customer",
	},
	models.User{
		Username: "johndoe",
		Email:    "johndoe@gmail.com",
		Password: "password",
		UserType: "customer",
	},
	models.User{
		Username: "jessica",
		Email:    "jessica@gmail.com",
		Password: "password",
		UserType: "customer",
	},
}

var products = []models.Product{
	models.Product{
		DisplayName: "Qtela Tempe Original 55 gr",
		Category:    "Makanan Ringan",
		Price:       6500,
		Quantity:    100,
		Description: "Makanan Ringan cocok untuk bersantai",
	},
	models.Product{
		DisplayName: "Chitato Sapi Bumbu Bakar 68 Gr",
		Category:    "Makanan Ringan",
		Price:       21000,
		Quantity:    200,
		Description: "Makanan Ringan cocok untuk bersantai",
	},
	models.Product{
		DisplayName: "Maxicorn Roasted Corn 160 Gr",
		Category:    "Makanan Ringan",
		Price:       10000,
		Quantity:    150,
		Description: "Makanan Ringan cocok untuk bersantai",
	},
}

var transactions = []models.Transaction{
	models.Transaction{
		TransactionId: "e7ee4de9-09ed-4566-9a7b-d76bf185fda4",
		CartId:        1,
		UserId:        2,
		ProductId:     1,
		Price:         30000,
		Total:         2,
		PaymentStatus: "Done",
	},
	models.Transaction{
		TransactionId: "e7ee4de9-09ed-4566-9a7b-d76bf185fda4",
		CartId:        2,
		UserId:        3,
		ProductId:     3,
		Price:         16000,
		Total:         2,
		PaymentStatus: "Done",
	},
}

var carts = []models.Cart{
	models.Cart{
		CartId:    1,
		ProductId: 1,
		UserId:    1,
		Total:     10,
	},
	models.Cart{
		CartId:    2,
		ProductId: 3,
		UserId:    1,
		Total:     8,
	},
	models.Cart{
		CartId:    3,
		ProductId: 3,
		UserId:    1,
		Total:     2,
	},
}

func Load(db *gorm.DB) {

	err := db.Debug().DropTableIfExists(&models.User{}, &models.Product{}, &models.Cart{}, &models.Transaction{}).Error
	if err != nil {
		log.Fatalf("cannot drop table: %v", err)
	}
	err = db.Debug().AutoMigrate(&models.User{}, &models.Product{}, &models.Cart{}, &models.Transaction{}).Error
	if err != nil {
		log.Fatalf("cannot migrate table: %v", err)
	}

	for i, _ := range users {
		err = db.Debug().Model(&models.User{}).Create(&users[i]).Error
		if err != nil {
			log.Fatalf("cannot seed users table: %v", err)
		}
	}

	for i, _ := range products {
		err = db.Debug().Model(&models.Product{}).Create(&products[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}
	}

	for i, _ := range carts {
		err = db.Debug().Model(&models.Cart{}).Create(&carts[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}
	}

	for i, _ := range transactions {
		err = db.Debug().Model(&models.Transaction{}).Create(&transactions[i]).Error
		if err != nil {
			log.Fatalf("cannot seed products table: %v", err)
		}
	}
}
