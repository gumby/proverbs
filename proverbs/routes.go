package proverbs

import (
	"net/http"

	"github.com/gorilla/mux"
)

type pathReader interface {
	read(name string, r *http.Request) string
}

type varReader struct{}

func (m varReader) read(name string, r *http.Request) string {
	return mux.Vars(r)[name]
}

func (s *Server) routes() {
	s.router.HandleFunc("/proverbs", s.handleProverbsGetAll()).Methods(http.MethodGet)
	s.router.HandleFunc("/proverbs/{id}", s.handleProverbsGet(varReader{})).Methods(http.MethodGet)

	s.router.MethodNotAllowedHandler = s.handleMethodNotAllowed()
}
