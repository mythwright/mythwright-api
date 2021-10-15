package api

import (
	"context"
	"fmt"
	"github.com/Zyian/mythwright-api/db"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	s  *http.Server
	db *db.Database
}

func NewServer(ctx context.Context) *Server {
	s := &Server{
		s: &http.Server{
			Addr:         viper.GetString("addr"),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
	s.s.Handler = buildMuxRouter(s, mux.NewRouter())

	database, err := db.NewDatabase(ctx)
	if err != nil {
		panic(err)
	}
	s.db = database

	return s
}

func buildMuxRouter(s *Server, r *mux.Router) *mux.Router {
	h := map[string]map[string]http.HandlerFunc{
		"POST": {
			`/items/inputs`:   postInputsHandler(s),
			`/items/mappings`: postItemMappings(s),
		},
	}
	for m, hm := range h {
		for path, handler := range hm {
			r.Methods(m).Path(path).HandlerFunc(handler)
		}
	}
	return r
}

func (s *Server) ListenAndServe(_ context.Context) error {
	return s.s.ListenAndServe()
}

func (s *Server) ListenAndServeTLS(_ context.Context) error {
	c := viper.Sub("tls")
	if c == nil {
		return fmt.Errorf("called listen tls but no config found")
	}
	return s.s.ListenAndServeTLS(c.GetString("cert_file"), c.GetString("key_file"))
}

func (s *Server) Shutdown(ctx context.Context) error {
	if err := s.s.Shutdown(ctx); err != nil {
		return err
	}

	return nil
}
