package controllers

import (
	"golang-thrive/api/models"
	"golang-thrive/api/responses"
	"net/http"

	"github.com/gorilla/mux"
)

func (server *Server) GetProducts(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}

	products, err := product.FindAllProduct(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, products)
}

func (server *Server) SearchProductByName(w http.ResponseWriter, r *http.Request) {
	product := models.Product{}
	vars := mux.Vars(r)
	products, err := product.FindProductByName(server.DB, vars["keyword"])
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, products)
}
