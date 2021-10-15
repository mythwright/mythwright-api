package api

import (
	"encoding/json"
	"github.com/Zyian/mythwright-api/db"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

func postItemMappings(s *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.WithFields(logrus.Fields{}).Info("got it mappings")
		if r.Body == nil || r.ContentLength < 1 {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte("no body provided"))
			return
		}

		var im db.ItemMapping
		if err := json.NewDecoder(r.Body).Decode(&im); err != nil {
			logrus.WithFields(logrus.Fields{"err": err}).Error("unable to decode request")
		}
		im.ID = primitive.NewObjectID()
		logrus.WithFields(logrus.Fields{"mapp": im}).Info("got it mappings")
	}
}
