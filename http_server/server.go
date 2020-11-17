package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// Player has a name and the number of wins
type Player struct {
	Name string
	Wins int
}

// PlayerStore stores information about players.
type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

// PlayerServer is an HTTP interface for player info.
type PlayerServer struct {
	store        PlayerStore
	http.Handler // embedding: https://golang.org/doc/effective_go.html#embedding
}

// NewPlayerServer returns a new PlayerServer
func NewPlayerServer(store PlayerStore) *PlayerServer {
	router := http.NewServeMux()

	p := &PlayerServer{
		store,
		router,
	}

	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))

	return p
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(p.store.GetLeague())
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}
