package main

import (
	"thrive-project/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connection() (*gorm.DB, error) {
	dsn := "host=localhost port=5432 user=postgres dbname=marketplace_api password=password sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	println(err)

	// var user User
	// var users []User

	// 	db.First[&user]

	// result := map[string]interfave{}{}
	// db.model(&User{}).First(&result)

	// 	db, err := gorm.Open(postgres.Open("host=localhost port=5432 user=postgres dbname=postgres password=postgres sslmode=disable"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&model.User{})
	return db, nil
}
