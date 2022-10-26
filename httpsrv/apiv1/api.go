package controller

import "thrive-project/vendor/gorm.io/gorm"

type Server struct {
	DB     *gorm.DB
	Router *mux.Router
}

