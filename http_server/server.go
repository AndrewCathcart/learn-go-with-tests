package main

import (
	"fmt"
	"net/http"
	"strings"
)

// PlayerServer is our server for player scores
func PlayerServer(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprint(w, GetPlayerScore(player))
}

// GetPlayerScore returns a players score
func GetPlayerScore(name string) string {
	if name == "Andrew" {
		return "20"
	}

	if name == "Rosie" {
		return "10"
	}

	return ""
}
