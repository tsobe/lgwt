package context

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

type SpyStore struct {
	response  string
	cancelled bool
	t         *testing.T
}

func (s *SpyStore) Fetch() string {
	time.Sleep(100 * time.Millisecond)
	return s.response
}

func (s *SpyStore) Cancel() {
	s.cancelled = true
}

func (s *SpyStore) assertNotCancelled() {
	s.t.Helper()
	if s.cancelled {
		s.t.Errorf("Store was unexpectedly cancelled")
	}
}

func (s *SpyStore) assertCancelled() {
	s.t.Helper()
	if !s.cancelled {
		s.t.Errorf("Store was not told to cancel")
	}
}

func TestServer(t *testing.T) {
	t.Run("Serves response", func(t *testing.T) {
		data := "Hello world"
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)

		if res.Body.String() != data {
			t.Errorf("Got %q, expected %q", res.Body.String(), data)
		}
		store.assertNotCancelled()
	})

	t.Run("Tells store to cancel work if request is cancelled", func(t *testing.T) {
		data := "Hello world"
		store := &SpyStore{response: data, t: t}
		srv := Server(store)

		req := httptest.NewRequest(http.MethodGet, "/", nil)

		cancellingCtx, cancel := context.WithCancel(req.Context())
		time.AfterFunc(5*time.Millisecond, cancel)
		req = req.WithContext(cancellingCtx)
		res := httptest.NewRecorder()

		srv.ServeHTTP(res, req)

		store.assertCancelled()
	})
}

type Store interface {
	Fetch() string
	Cancel()
}

func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		data := make(chan string, 1)

		go func() {
			data <- store.Fetch()
		}()

		select {
		case storeData := <-data:
			_, _ = fmt.Fprint(w, storeData)
		case <-ctx.Done():
			store.Cancel()
		}
	}
}
