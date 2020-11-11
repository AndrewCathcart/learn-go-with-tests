package sync

import "sync"

// Make a counter which is safe to use concurrently

// Counter holds a value and a mutex
type Counter struct {
	mu    sync.Mutex
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
