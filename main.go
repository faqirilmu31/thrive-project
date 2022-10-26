package main

import (
	"log"
	"net/http"

	"thrive-project/controller"
	"github.com/labstack/echo/v4"
)

func main() {
	// initialize echo framework
	e := echo.New()

	// connect to database
	db, err := Connection()
	

	// _, err := Connection()
	if err != nil {
		log.Panic(err)
		return
	}

	h := controller.New(db)


	// root
	e.GET("/", func(c echo.Context) error {
		result := map[string]string{
			"response_code": "200",
			"message":       "Success to connect service",
		}

		return c.JSON(http.StatusOK, result)
	})

	cust := e.Group("customer")
	cust.GET("", h.GetCustomer)
	// cust.GET("/:id", h.GetCustomerByID)
	// cust.POST("", h.CreateCustomer)
	// cust.PUT("/:id", h.UpdateCustomer)
	// cust.DELETE("/:id", h.DeleteCustomer)

	// start service echo
	e.Logger.Fatal(e.Start(":5002"))
}