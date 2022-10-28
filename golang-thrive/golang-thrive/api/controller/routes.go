package controllers

import "golang-thrive/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Routes
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Routes
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users Routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Product Routes
	s.Router.HandleFunc("/product/list", middlewares.SetMiddlewareJSON(s.GetProducts)).Methods("GET")
	s.Router.HandleFunc("/product/search/{keyword}", middlewares.SetMiddlewareJSON(s.SearchProductByName)).Methods("GET")

	//Cart Routes
	s.Router.HandleFunc("/cart/add", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.AddToCart))).Methods("POST")
	s.Router.HandleFunc("/cart/remove", middlewares.SetMiddlewareJSON(s.DeleteCart)).Methods("POST")

	//Transaction Routes
	s.Router.HandleFunc("/transaction/checkOut", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CheckOut))).Methods("POST")
	s.Router.HandleFunc("/transaction/checkTotalPrice", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.CheckTotalPrice))).Methods("GET")
}
