package maths

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func secondHandPoint(seconds time.Time) Point {
	return angleToPoint(secondsInRadians(seconds))
}

func secondsInRadians(seconds time.Time) float64 {
	return math.Pi / (30 / float64(seconds.Second()))
}

func minutesInRadians(minutes time.Time) float64 {
	return (secondsInRadians(minutes) / 60) + (math.Pi / (30 / float64(minutes.Minute())))
}

func minuteHandPoint(time time.Time) Point {
	return angleToPoint(minutesInRadians(time))
}

func angleToPoint(angle float64) Point {
	return Point{math.Sin(angle), math.Cos(angle)}
}
