package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"golang-thrive/api/auth"
	"golang-thrive/api/models"
	"golang-thrive/api/responses"
	"golang-thrive/api/utils/formaterror"
	"io/ioutil"
	"net/http"
)

func (server *Server) AddToCart(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}
	user_id, err := auth.ExtractTokenID(r)
	if err != nil {
		responses.ERROR(w, http.StatusUnauthorized, err)
	}

	cart := models.Cart{}
	cart.UserId = user_id
	err = json.Unmarshal(body, &cart)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	if cart.Total == 0 {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = cart.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	cartAdded, err := cart.AddCart(server.DB)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, cartAdded.CartId))
	responses.JSON(w, http.StatusCreated, cartAdded)
}

func (server *Server) DeleteCart(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	var cartId cartID
	err = json.Unmarshal(body, &cartId)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	user_id, err := auth.ExtractTokenID(r)
	if err != nil {
		fmt.Println(err)
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	cart := models.Cart{}
	carts, err := cart.FindCartByID(server.DB, cartId.CartId)
	if err != nil {
		responses.ERROR(w, http.StatusNotFound, err)
		return
	}

	if user_id != carts.UserId {
		responses.ERROR(w, http.StatusUnauthorized, errors.New("Unauthorized"))
		return
	}

	err = cart.DeleteCart(server.DB, carts.CartId)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", cartId.CartId))
	responses.JSON(w, http.StatusNoContent, "")
}
