package main

import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

func TestCountdown(t *testing.T) {
	t.Run("Countdown from 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		want := `3
2
1
Go!`

		Countdown(buffer, &SpyCountdownOperations{})

		got := buffer.String()

		if got != want {
			t.Errorf("Got %s, expected %s", got, want)
		}
	})

	t.Run("Sleep before every print", func(t *testing.T) {
		expectedCallOrder := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}
		spyCountdownOperations := &SpyCountdownOperations{}

		Countdown(spyCountdownOperations, spyCountdownOperations)

		if !reflect.DeepEqual(spyCountdownOperations.Calls, expectedCallOrder) {
			t.Errorf(
				"Expected calls in following order %v, got %v",
				expectedCallOrder,
				spyCountdownOperations.Calls,
			)
		}
	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := ConfigurableSleeper{sleepTime, spyTime.Sleep}

	sleeper.Sleep()

	if spyTime.DurationSlept != sleepTime {
		t.Errorf("Should have slept for %s, but slept for %s", sleepTime, spyTime.DurationSlept)
	}
}

const sleep = "s"
const write = "w"

type SpyCountdownOperations struct {
	Calls []string
}

func (s *SpyCountdownOperations) Write(b []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

type SpyTime struct {
	DurationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.DurationSlept = duration
}
