package sync

import "sync"

// Make a counter which is safe to use concurrently

// Counter holds a value and a mutex
type Counter struct {
	mu    sync.Mutex // We don't want to expose this to the public API
	value int
}

// NewCounter creates a new counter
func NewCounter() *Counter {
	return &Counter{}
}

// Inc increases the counter by 1
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++
}

// Value returns the current value of the counter
func (c *Counter) Value() int {
	return c.value
}

// Locks over channels and goroutines
// https://github.com/golang/go/wiki/MutexOrChannel
// Use channels when passing ownership of data
// Use mutexes for managing state

// Remember to use go vet in build scripts as it can alert to subtle bugs in our code
