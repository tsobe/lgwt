package maths

import (
	"math"
	"time"
)

const (
	secondsInHalfClock = 30
	minutesInHalfClock = 30
	minutesInClock     = 2 * minutesInHalfClock
	hoursInHalfClock   = 6
	hoursInClock       = 2 * hoursInHalfClock
)

type Point struct {
	X float64
	Y float64
}

func secondHandPoint(seconds time.Time) Point {
	return angleToPoint(secondsInRadians(seconds))
}

func secondsInRadians(seconds time.Time) float64 {
	return math.Pi / (secondsInHalfClock / float64(seconds.Second()))
}

func minutesInRadians(minutes time.Time) float64 {
	return (secondsInRadians(minutes) / minutesInClock) + (math.Pi / (minutesInHalfClock / float64(minutes.Minute())))
}

func hoursInRadians(hours time.Time) float64 {
	return (minutesInRadians(hours) / hoursInClock) + (math.Pi / (hoursInHalfClock / float64(hours.Hour()%hoursInClock)))
}

func minuteHandPoint(time time.Time) Point {
	return angleToPoint(minutesInRadians(time))
}

func hourHandPoint(time time.Time) Point {
	return angleToPoint(hoursInRadians(time))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}
