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

// Covered how to test a HTTP handler that has had the requets cancelled by the client
// How to use context to manage cancellation
// How to write a function that accepts context and uses it to cancel itself by using goroutines, select and channels
// Follow Google's guidelines as to how to manage cancellation by propagating requets scoped context through the call-stack
// How to create our own spy for http.ResponseWriter if we ever need it

// Don't use context.Value - it's an untpyed map, no type-safety and you have to handle it not actually containing our value
// https://faiface.github.io/post/context-should-go-away-go2/
// https://blog.golang.org/context
