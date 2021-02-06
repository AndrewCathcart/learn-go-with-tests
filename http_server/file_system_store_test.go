package main

import (
	"strings"
	"testing"
)

func TestFileSystemStore(t *testing.T) {
	t.Run("/league from a reader", func(t *testing.T) {
		database := strings.NewReader(`[
			{"Name": "Andy", "Wins": 100},
			{"Name": "Rosie", "Wins": 1}]`)

		store := FileSystemPlayerStore{database}

		received := store.GetLeague()
		expected := []Player{
			{"Andy", 100},
			{"Rosie", 1},
		}

		assertLeague(t, received, expected)
	})

}
