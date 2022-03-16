package context

import (
	"context"
	"errors"
	"log"
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

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", ctx.Err()
	case res := <-data:
		return res, nil
	}
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

type SpyResponseWriter struct {
	written bool
}

func (w *SpyResponseWriter) Header() http.Header {
	w.written = true
	return nil
}

func (w *SpyResponseWriter) Write(bytes []byte) (int, error) {
	w.written = true
	return 0, errors.New("not implemented")
}

func (w *SpyResponseWriter) WriteHeader(statusCode int) {
	w.written = true
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
		res := &SpyResponseWriter{}

		srv.ServeHTTP(res, req)

		if res.written {
			t.Errorf("A response shouldn't have written")
		}
	})
}
