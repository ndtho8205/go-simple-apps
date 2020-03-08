package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   []Player
}

func (s *StubPlayerStore) GetPlayerScore(name string) (int, bool) {
	score, ok := s.scores[name]
	return score, ok
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeagues() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		},
		[]string{},
		nil,
	}
	server := NewPlayerServer(&store)

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
		server := NewPlayerServer(&store)

		req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)
	})

	t.Run("record player when POST", func(t *testing.T) {
		store := StubPlayerStore{}
		server := NewPlayerServer(&store)

		req, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusAccepted)

		got := len(store.winCalls)
		want := 1
		if got != want {
			t.Fatalf("got %d calls to RecordWin want %d", got, want)
		}

		if store.winCalls[0] != "Pepper" {
			t.Errorf("wrong player got %q want %q", store.winCalls[0], "Pepper")
		}
	})
}

func TestLeague(t *testing.T) {

	t.Run("return 200", func(t *testing.T) {
		wantedLeagues := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{nil, nil, wantedLeagues}
		server := NewPlayerServer(&store)

		req := newGetLeaguesRequest()
		res := httptest.NewRecorder()

		server.ServeHTTP(res, req)

		assertStatus(t, res.Code, http.StatusOK)

		assertContentType(t, res, "application/json")

		got := getLeaguesFromResponse(t, res)
		if !reflect.DeepEqual(got, wantedLeagues) {
			t.Errorf("got %v want %v", got, wantedLeagues)
		}
	})
}

func newGetPlayerScoreRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newPostWinRequest(player string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", player), nil)
	return req
}

func newGetLeaguesRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}
func getLeaguesFromResponse(t *testing.T, res *httptest.ResponseRecorder) []Player {
	var got []Player
	err := json.NewDecoder(res.Body).Decode(&got)
	if err != nil {
		t.Fatalf("unable to parse response %q, error: %v", res.Body, err)
	}

	return got
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()

	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func assertContentType(t *testing.T, res *httptest.ResponseRecorder, want string) {
	t.Helper()

	if res.Result().Header.Get("content-type") != "application/json" {
		t.Fatalf("response did not have content-type of application/json, got %v", res.Result().Header.Get("content-type"))
	}

}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
