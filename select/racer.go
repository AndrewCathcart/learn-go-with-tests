package main

import (
	"net/http"
)

// Racer takes two urls, and compares the Get time of each, returning the fastest
func Racer(a, b string) (winner string, error error) {
	// select lets you wait on multiple channels, the first channel to send a value has it's case statement executed
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
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
