package maths

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const clockCenterX = 150
const clockCenterY = 150

func SecondHand(time time.Time) Point {
	p := secondHandPoint(time)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate
	return p
}

func secondHandPoint(seconds time.Time) Point {
	angle := secondsInRadians(seconds)
	return Point{math.Sin(angle), math.Cos(angle)}
}

func secondsInRadians(seconds time.Time) float64 {
	return math.Pi / (30 / float64(seconds.Second()))
}
