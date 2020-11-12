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

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Andrew": 20,
			"Rosie":  10,
		},
		nil,
	}
	server := &PlayerServer{&store}

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
	}
	server := &PlayerServer{&store}

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

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	server := &PlayerServer{NewInMemoryPlayerStore()}
	player := "Andy"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	response := httptest.NewRecorder()
	server.ServeHTTP(response, newGetScoreRequest(player))
	assertStatus(t, response.Code, http.StatusOK)

	assertResponseBody(t, response.Body.String(), "2")
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
