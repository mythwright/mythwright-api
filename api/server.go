package api

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"net/http"
	"time"
)

func init() {
	viper.SetDefault("addr", ":8000")
}

type Server struct {
	s  *http.Server
	db *mongo.Client
}

func NewServer(ctx context.Context) *Server {
	r := buildMuxRouter(mux.NewRouter())

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(viper.GetString("db_uri")))
	if err != nil {
		logrus.Panic(err)
	}

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		logrus.Panic(err)
	}
	logrus.Info("Successfully Connected and Pinged MongoDB")

	return &Server{
		s: &http.Server{
			Handler:      r,
			Addr:         viper.GetString("addr"),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
	}
}

func buildMuxRouter(r *mux.Router) *mux.Router {
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

func (s *Server) PingDB(ctx context.Context) error {
	return s.db.Ping(ctx, readpref.Primary())
}
