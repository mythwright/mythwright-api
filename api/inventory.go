package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (s *Server) LoadItems() {

}

func postInputsHandler(s *Server) func(w http.ResponseWriter, r *http.Request) {
	logrus.WithFields(logrus.Fields{}).Info("got it")
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) postInputs(ctx context.Context) {

}
