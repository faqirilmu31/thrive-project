package controller

import (
	"thrive-project/model"
	"net/http"
	"log"
	"os"
	"github.com/labstack/echo/v4"
)

func (h Handler) GetUser(c echo.Context) error {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	log.SetOutput(file)

	log.Println("Masuk ke get all user")
	var result []model.User
	var customer model.User
	h.db.First(&customer)
	result = append(result, customer)

	log.Println("Berhasil get all user")
	return c.JSON(http.StatusOK, result)
}