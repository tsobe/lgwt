package main

import (
	"bytes"
	"testing"
)

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func TestCountdown(t *testing.T) {
	buffer := &bytes.Buffer{}
	sleeper := &SpySleeper{}
	want := `3
2
1
Go!`

	Countdown(buffer, sleeper)

	got := buffer.String()

	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}

	if sleeper.Calls != 4 {
		t.Errorf("Expected %d calls, got %d", 4, sleeper.Calls)
	}
}
