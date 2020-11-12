package context

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"
)

// SpyStore lets you simulate a Store
type SpyStore struct {
	response string
	t        *testing.T
}

// Fetch is a mock fetch function
func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	// Simulate a super slow process by sleeping and concatenating each char of the response
	go func() {
		var result string
		for _, c := range s.response {
			select {
			case <-ctx.Done():
				s.t.Log("spy store was cancelled")
				return
			default:
				time.Sleep(10 * time.Millisecond)
				result += string(c)
			}
		}
		data <- result
	}()

	// Using select with two channels to get the first one that we get a result for
	select {
	case <-ctx.Done(): // the channel that .Done() returns when work should be cancelle
		return "", ctx.Err()
	case res := <-data: // the actual response
		return res, nil
	}
}

// SpyResponseWriter checks whether a response has been written.
type SpyResponseWriter struct {
	written bool
}

// Header will mark written to true.
func (s *SpyResponseWriter) Header() http.Header {
	s.written = true
	return nil
}

// Write will mark written to true.
func (s *SpyResponseWriter) Write([]byte) (int, error) {
	s.written = true
	return 0, errors.New("not implemented")
}

// WriteHeader will mark written to true.
func (s *SpyResponseWriter) WriteHeader(statusCode int) {
	s.written = true
}
