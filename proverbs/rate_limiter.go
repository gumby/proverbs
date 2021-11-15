package proverbs

import (
	"net/http"
	"sync"
	"time"
)

type limitStore map[string]*fixedLimiter
type fixedLimiter struct {
	mu     sync.Mutex
	tokens uint
}

func newFixedLimiter(max uint, refill uint, d time.Duration) *fixedLimiter {
	t := &fixedLimiter{
		tokens: max,
	}

	ticker := time.NewTicker(d)
	go func() {
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				new := t.tokens + refill
				if new > max {
					new = max
				}
				t.mu.Lock()
				t.tokens = new
				t.mu.Unlock()
			}
		}
	}()
	return t
}

func (t *fixedLimiter) getToken() bool {
	if t.tokens <= 0 {
		return false
	}
	t.mu.Lock()
	defer t.mu.Unlock()

	t.tokens--
	return true
}

func (s *Server) throttle(max uint, refill uint, d time.Duration) func(http.Handler) http.HandlerFunc {
	var store limitStore = make(limitStore)
	return func(h http.Handler) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// This is a hack replacement of a login/authn/api key
			userID := r.Header.Get("X-USER")
			if userID == "" {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
			t, ok := store[userID]
			if !ok {
				t = newFixedLimiter(max, refill, d)
				store[userID] = t
			}
			if !t.getToken() {
				w.WriteHeader(http.StatusTooManyRequests)
				return
			}
			h.ServeHTTP(w, r)
		}
	}
}
