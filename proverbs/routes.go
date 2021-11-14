package proverbs

import (
	"net/http"
	"time"

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
	s.router.HandleFunc("/proverbs", s.throttle(25, 25, 15*time.Second)(s.handleProverbsGetAll())).Methods(http.MethodGet)
	s.router.HandleFunc("/proverbs/{id}", s.handleProverbsGet(varReader{})).Methods(http.MethodGet)

	s.router.MethodNotAllowedHandler = s.handleMethodNotAllowed()
}
