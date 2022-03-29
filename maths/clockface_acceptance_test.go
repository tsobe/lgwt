package maths

import (
	"testing"
)

func TestSecondHandAtMidnight(t *testing.T) {
	tm := simpleTime(1, 0, 0)

	want := Point{X: clockCenterX, Y: clockCenterY - secondHandLength}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	tm := simpleTime(1, 0, 30)

	want := Point{X: clockCenterX, Y: clockCenterY + secondHandLength}
	got := SecondHand(tm)

	if got != want {
		t.Errorf("Got %v, wanted %v", got, want)
	}
}
