package controllers

import (
	"golang-thrive/api/responses"
	"net/http"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To THRIVE: Telkomsel Tech Enthusiast Bootcamp")

}
