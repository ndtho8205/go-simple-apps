package context

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	t         *testing.T
	response  string
	cancelled bool
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func TestServer(t *testing.T) {
	t.Run("fetch successfully", func(t *testing.T) {
		data := "Hello, World"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		store.assertWasNotCancelled()

		if res.Body.String() != data {
			t.Errorf("got %q want %q", res.Body.String(), data)
		}
	})

	t.Run("cancel while fetching", func(t *testing.T) {
		data := "Hello, World"
		store := &SpyStore{response: data, t: t}
		server := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)

		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		store.assertWasCancelled()
	})
}

func (s *SpyStore) assertWasCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("it should be cancelled")
	}

}

func (s *SpyStore) assertWasNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("it should not be cancelled")
	}
}
