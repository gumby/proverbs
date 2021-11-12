package proverbs

import (
	"fmt"
	"net/http"
	"strconv"
)

func (s *Server) handleProverbsGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		proverbs, err := s.store.GetAll()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		for _, p := range proverbs {
			fmt.Fprintf(w, "%s\n", p)
		}
	}
}

func (s *Server) handleProverbsGet(pr pathReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sID := pr.read("id", r)
		if sID == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		id, err := strconv.Atoi(sID)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		p, err := s.store.Get(id)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "%s\n", p)
	}
}

func (s *Server) handleMethodNotAllowed() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}
