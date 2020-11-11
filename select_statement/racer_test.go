package select_statement

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {
	t.Run("returns the fastest response", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		defer slowServer.Close()
		fastServer := makeDelayedServer(0 * time.Millisecond)
		defer fastServer.Close()

		expected := fastServer.URL
		actual, err := Racer(slowServer.URL, fastServer.URL)

		if err != nil {
			t.Fatalf("did not expect an error but goe one %v", err)
		}

		if expected != actual {
			t.Errorf("got %q, want %q", actual, expected)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		racerTimeout := 1 * time.Millisecond
		serverResponseDelay := 2 * racerTimeout

		serverA := makeDelayedServer(serverResponseDelay)
		defer serverA.Close()

		_, err := ConfigurableRacer(serverA.URL, serverA.URL, racerTimeout)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
