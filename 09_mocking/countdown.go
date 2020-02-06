package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countdownStart = 3
	finalWord      = "Go!"
	delay          = 1
)

type Sleeper interface {
	Sleep()
}

func main() {
	Countdown(os.Stdout, &ConfigurableSleeper{1 * time.Second, time.Sleep})
}

type DefaultSleeper struct{}

func (sleeper DefaultSleeper) Sleep() {
	time.Sleep(delay * time.Second)
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (sleeper *ConfigurableSleeper) Sleep() {
	sleeper.sleep(sleeper.duration)
}

func Countdown(w io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(w, i)
	}
	sleeper.Sleep()
	fmt.Fprint(w, finalWord)
}
