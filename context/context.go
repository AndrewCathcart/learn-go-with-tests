package context

import (
	"context"
	"fmt"
	"net/http"
)

// How do we handle long running resource-intensive applications (sometimes in goroutines) if the action that caused them was cancelled / failed?
// Need to handle these processes in a consistent way through our app. This is done through the context package.

// Store represents a data store
type Store interface {
	Fetch(ctx context.Context) (string, error)
}

// Server represents our server
func Server(store Store) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := store.Fetch(r.Context())

		if err != nil {
			return
		}
		fmt.Fprint(w, data)
	}
}
