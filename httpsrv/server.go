package httpsrv

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Server struct {
	ListenIP   string
	ListenPort string
	DBConnStr  string

	router chi.Router
	DB     *gorm.DB
}

func (s *Server) Init() error {
	var err error
	db, err := gorm.Open(postgres.Open(s.DBConnStr), &gorm.Config{})
	if err != nil {
		return err
	}
	s.DB = db

	s.router = chi.NewRouter()
	s.router.Route("/api", func(r chi.Router) {
	})

	return nil
}

func (s *Server) Start() error {
	listenAddr := fmt.Sprintf("%s:%s", s.ListenIP, s.ListenPort)
	return http.ListenAndServe(listenAddr, s.router)
}