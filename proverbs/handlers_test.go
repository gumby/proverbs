package proverbs

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type mockStore struct {
	size int
}

func (m mockStore) GetAll() ([]Proverb, error) {
	return []Proverb{
		"Proverb 1",
		"Proverb 2",
	}, nil
}

type mockPathReader struct {
	val string
}

func (m mockPathReader) read(name string, r *http.Request) string {
	return m.val
}

func (m mockStore) Get(id int) (Proverb, error) {
	return Proverb(fmt.Sprintf("Proverb %d", id)), nil
}

func TestProverbsGetAll(t *testing.T) {
	srv := NewServer(mockStore{})
	r := httptest.NewRequest(http.MethodGet, "/proverbs", nil)
	w := httptest.NewRecorder()
	srv.handleProverbsGetAll().ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected 200 but got %d", w.Code)
	}
}

func TestProverbGet(t *testing.T) {
	srv := NewServer(mockStore{})
	r := httptest.NewRequest(http.MethodGet, "/proverbs/1", nil)
	w := httptest.NewRecorder()
	srv.handleProverbsGet(mockPathReader{val: "1"}).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 but got %d", w.Code)
	}

	body, _ := ioutil.ReadAll(w.Body)
	if !strings.Contains(string(body), "Proverb 1") {
		t.Fatalf("expected string containing Proverb 1 but got %s", string(body))
	}
}
