package maths

import (
	"fmt"
	"io"
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

const secondHandLength = 90
const minuteHandLength = 80
const clockCenterX = 150
const clockCenterY = 150

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`
const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`
const svgEnd = `</svg>`

func SecondHand(time time.Time) Point {
	p := secondHandPoint(time)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate
	return p
}

func SVGWriter(w io.Writer, t time.Time) {
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	writeSecondHand(w, t)
	writeMinuteHand(w, t)
	io.WriteString(w, svgEnd)
}

func writeSecondHand(w io.Writer, t time.Time) {
	p := makeHand(secondHandPoint(t), secondHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}

func writeMinuteHand(w io.Writer, t time.Time) {
	p := makeHand(minuteHandPoint(t), minuteHandLength)
	fmt.Fprintf(w, `<line x1="150" y1="150" x2="%.3f" y2="%.3f" style="fill:none;stroke:#000;stroke-width:3px;"/>`, p.X, p.Y)
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X, -p.Y}
	return Point{p.X + clockCenterX, p.Y + clockCenterY}
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
