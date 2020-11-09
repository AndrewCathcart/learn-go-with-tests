package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

// Sleeper interface with a Sleep function
type Sleeper interface {
	Sleep()
}

// ConfigurableSleeper represents a Sleeper that can have it's sleep time changed
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

// Sleep for a duration
func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

const finalWord = "Go!"
const countdownStart = 3

// Countdown prints 3, 2, 1, Go!
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
