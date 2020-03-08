package main

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetPlayerScoreRequest(player))

		assertStatus(t, res.Code, http.StatusOK)
		assertResponseBody(t, res.Body.String(), "3")
	})

	t.Run("get leagues", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetLeaguesRequest())

		assertStatus(t, res.Code, http.StatusOK)
		assertContentType(t, res, "application/json")

		got := getLeaguesFromResponse(t, res)
		want := []Player{
			{"Pepper", 3},
		}
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
