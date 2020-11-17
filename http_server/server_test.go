package main

import (
	"encoding/json"
	"fmt"
	"io"
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

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() []Player {
	return s.league
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Andrew": 20,
			"Rosie":  10,
		},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("returns Andrew's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Andrew", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns Rosie's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Rosie", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("returns a 404 if the player doesn't exist", func(t *testing.T) {
		request := newGetScoreRequest("Appa")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{},
		nil,
		nil,
	}
	server := NewPlayerServer(&store)

	t.Run("it returns wins on POST", func(t *testing.T) {
		player := "Andrew"

		request := newPostWinRequest("Andrew")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)

		if len(store.winCalls) != 1 {
			t.Errorf("received %d calls to RecordWin, expected %d", len(store.winCalls), 1)
		}

		if store.winCalls[0] != player {
			t.Errorf("did not store correct winner, received %q but expected %q", store.winCalls[0], player)
		}
	})
}

func TestLeagueTable(t *testing.T) {
	t.Run("it returns the league table as JSON", func(t *testing.T) {
		expectedLeague := []Player{
			{"Andy", 1},
			{"Rosie", 2},
			{"Appa", 3},
		}
		store := StubPlayerStore{nil, nil, expectedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		received := getLeagueFromResponse(t, response.Body)

		assertStatus(t, response.Code, http.StatusOK)
		assertLeague(t, received, expectedLeague)
		assertContentType(t, response, jsonContentType)
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertResponseBody(t *testing.T, received, expected string) {
	t.Helper()
	if received != expected {
		t.Errorf("response body is wrong, received %q expected %q", received, expected)
	}
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)

	if err != nil {
		t.Fatalf("Unable to parse response from server %q into slice of Player, '%v'", body, err)
	}

	return
}

func assertLeague(t *testing.T, got, want []Player) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

const jsonContentType = "application/json"

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Result().Header.Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.Result().Header)
	}
}
