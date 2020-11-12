package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type StubPlayerStore struct {
	scores map[string]int
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func TestGETPlayers(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{
			"Andrew": 20,
			"Rosie":  10,
		},
	}
	server := &PlayerServer{&store}

	t.Run("returns Andrew's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Andrew", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "20")
	})

	t.Run("returns Rosie's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Rosie", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertResponseBody(t, response.Body.String(), "10")
	})
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
