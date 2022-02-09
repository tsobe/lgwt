package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("Fastest wins", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL
		want := fastUrl

		got, err := Racer(slowUrl, fastUrl)

		if err != nil {
			t.Errorf("Expected no error, got %q", err)
		}

		if got != want {
			t.Errorf("Got %s, expected %s", got, want)
		}
	})

	t.Run("Error is returned if timeout is exceeded", func(t *testing.T) {
		slowServer := makeDelayedServer(11 * time.Second)
		fastServer := makeDelayedServer(10 * time.Second)

		defer slowServer.Close()
		defer fastServer.Close()

		slowUrl := slowServer.URL
		fastUrl := fastServer.URL

		_, err := Racer(slowUrl, fastUrl)

		if err == nil {
			t.Error("Expected an error but didn't get one")
		}

	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
