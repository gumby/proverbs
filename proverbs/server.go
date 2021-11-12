package proverbs

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	router *mux.Router
	store  Store
}

func NewServer(store Store) *Server {
	s := &Server{
		router: mux.NewRouter(),
		store:  store,
	}
	s.routes()
	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
