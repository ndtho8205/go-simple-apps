package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("race between two URLs", func(t *testing.T) {
		slowServer := makeDelayServer(2 * time.Millisecond)
		fastServer := makeDelayServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one '%v'", err)
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("return an error when timeout", func(t *testing.T) {
		firstServer := makeDelayServer(24 * time.Millisecond)
		secondServer := makeDelayServer(35 * time.Millisecond)

		defer firstServer.Close()
		defer secondServer.Close()

		_, err := ConfigurableRacer(firstServer.URL, secondServer.URL, 20*time.Millisecond)
		if err == nil {
			t.Error("expected an error but didn't get one")
		}

	})
}

func makeDelayServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
