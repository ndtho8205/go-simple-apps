package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
	}
	server := &PlayerServer{&store}

	cases := []struct {
		name           string
		player         string
		wantHTTPStatus int
		wantScore      string
	}{
		{"Returns Pepper's score", "Pepper", http.StatusOK, "20"},
		{"Returns Floyd's score", "Floyd", http.StatusOK, "10"},
		{"Returns 404 on missing player", "Apollo", http.StatusNotFound, "0"},
	}

	for _, test := range cases {
		t.Run(test.name, func(t *testing.T) {
			req := newGetPlayerScoreRequest(test.player)
			res := httptest.NewRecorder()

			server.ServeHTTP(res, req)

			gotHTTPStatus := res.Code
			gotScore := res.Body.String()

			assertStatus(t, gotHTTPStatus, test.wantHTTPStatus)
			assertResponseBody(t, gotScore, test.wantScore)
		})
	}
}

func TestPostPlayer(t *testing.T) {
	t.Run("accepted on POST", func(t *testing.T) {
		store := StubPlayerStore{}
		server := PlayerServer{&store}

		req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)
	})

	t.Run("record player when POST", func(t *testing.T) {
		store := StubPlayerStore{}
		server := PlayerServer{&store}

		req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)

		got := len(store.winCalls)
		want := 1
		if got != want {
			t.Errorf("got %d calls to RecordWin want %d", got, want)
		}

		if store.winCalls[0] != "Pepper" {
			t.Errorf("wrong player got %q want %q", store.winCalls[0], "Pepper")
		}
	})
}

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryPlayerStore()
	server := PlayerServer{store}
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	res := httptest.NewRecorder()
	server.ServeHTTP(res, newGetPlayerScoreRequest(player))

	assertStatus(t, res.Code, http.StatusOK)
	assertResponseBody(t, res.Body.String(), "3")
}

func newGetPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
