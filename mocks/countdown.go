package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const finalWord = "Go!"

func main() {
	Countdown(os.Stdout, ConfigurableSleeper{1 * time.Second, time.Sleep})
}

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(out, i)
	}
	sleeper.Sleep()
	fmt.Fprint(out, finalWord)
}
