package api

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

type Server struct {
	s *http.Server
}

func NewServer() *Server {

	r := mux.NewRouter()

	srv := &http.Server{
		Handler:      r,
		Addr:         "0.0.0.0:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	s := &Server{
		s: srv,
	}

	return s
}
