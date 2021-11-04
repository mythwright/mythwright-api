package api

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/Zyian/mythwright-api/db"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"net/http"
	"time"
)

type Server struct {
	s           *http.Server
	db          *db.Database
	discordAuth *oauth2.Config
}

func NewServer(ctx context.Context) (*Server, error) {
	s := &Server{
		s: &http.Server{
			Addr:         viper.GetString("addr"),
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		},
		discordAuth: &oauth2.Config{
			ClientID:     viper.Sub("auth").Sub("discord").GetString("client_id"),
			ClientSecret: viper.Sub("auth").Sub("discord").GetString("client_secret"),
			Endpoint: oauth2.Endpoint{
				AuthURL:  "https://discord.com/api/oauth2/authorize",
				TokenURL: "https://discord.com/api/oauth2/token",
			},
			RedirectURL: viper.Sub("auth").Sub("discord").GetString("redirect_url"),
			Scopes:      []string{"identify"},
		},
	}
	s.s.Handler = buildMuxRouter(s, mux.NewRouter())

	database, err := db.NewDatabase(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to create db: (%v)", err)
	}
	s.db = database

	return s, nil
}

func buildMuxRouter(s *Server, r *mux.Router) *mux.Router {
	h := map[string]map[string]http.HandlerFunc{
		"POST": {
			`/items/inputs`:   postInputsHandler(s),
			`/items/mappings`: postItemMappings(s),
		},
		"GET": {
			"/auth/discord":          getDiscordAuth(s),
			"/auth/discord/callback": getDiscordAuthCallback(s),
			"/auth/gw2auth/callback": nil,
			"/auth/basic":            nil,
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

func (s *Server) AuthHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}

func getDiscordAuth(s *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		expire := time.Now().Add(1 * time.Hour)
		b := make([]byte, 16)
		_, _ = rand.Read(b)
		state := base64.URLEncoding.EncodeToString(b)
		cookie := &http.Cookie{Name: "auth.state", Value: state, Expires: expire}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, s.discordAuth.AuthCodeURL(state), http.StatusTemporaryRedirect)
	}
}

func getDiscordAuthCallback(_ *Server) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		authState, _ := r.Cookie("auth.state")

		if r.FormValue("state") != authState.Value {
			log.Info("invalid state")
			return
		}

		cookie := &http.Cookie{Name: "auth.code", Value: r.FormValue("code"), Expires: time.Now().Add(24 * time.Hour)}
		http.SetCookie(w, cookie)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
	}
}
