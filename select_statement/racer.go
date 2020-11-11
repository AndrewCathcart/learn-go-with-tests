package select_statement

import (
	"fmt"
	"net/http"
	"time"
)

var tenSecTimeout = 10 * time.Second

// Racer compares the response times of a and b, returning the fastest one, timing out after 10s.
func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecTimeout)
}

// ConfigurableRacer compares the response times of a and b, returning the fastest one, timing out after the specified timeout.
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	// select lets you wait on multiple channels, the first channel to send a value has it's case statement executed
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// chan struct{} is the smallest data type available from a memory perspective, so we get no allocation versus e.g. a boolean
	ch := make(chan struct{})

	// close the channel as soon as we get a result
	go func() {
		http.Get(url)
		close(ch)
	}()

	return ch
}
