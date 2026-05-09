package clockface

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

const (
	clockCentreX       = 150
	clockCentreY       = 150
	secondHandLength   = 90
	minuteHandLength   = 80
	hourHandLength     = 50
	secondsInHalfClock = 30
	minutesInHalfClock = 30
	hoursInHalfClock   = 6
)

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}

func makeHand(p Point, length float64) Point {
	p = Point{p.X * length, p.Y * length}
	p = Point{p.X + clockCentreX, -p.Y + clockCentreY}
	return p
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (secondsInHalfClock / float64(t.Second())))
}

func secondHandPoint(t time.Time) Point {
	angle := secondsInRadians(t)
	return angleToPoint(angle)
}

func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	return makeHand(p, secondHandLength)
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / (2 * minutesInHalfClock)) + (math.Pi / (minutesInHalfClock / float64(t.Minute())))
}

func minuteHandPoint(t time.Time) Point {
	angle := minutesInRadians(t)
	return angleToPoint(angle)
}

func MinuteHand(t time.Time) Point {
	p := minuteHandPoint(t)
	return makeHand(p, minuteHandLength)
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / (2 * hoursInHalfClock)) + (math.Pi / (hoursInHalfClock / float64(t.Hour()%12)))
}

func hourHandPoint(t time.Time) Point {
	angle := hoursInRadians(t)
	return angleToPoint(angle)
}

func HourHand(t time.Time) Point {
	p := hourHandPoint(t)
	return makeHand(p, hourHandLength)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`

func handTag(p Point, color string) string {
	return fmt.Sprintf(`<line x1="%d" y1="%d" x2="%.3f" y2="%.3f" style="fill:none;stroke:%s;stroke-width:3px;"/>`, clockCentreX, clockCentreY, p.X, p.Y, color)
}

func SVGWriter(w io.Writer, t time.Time) {
	sh := SecondHand(t)
	mh := MinuteHand(t)
	hh := HourHand(t)
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, handTag(sh, "#f00"))
	io.WriteString(w, handTag(mh, "#000"))
	io.WriteString(w, handTag(hh, "#000"))
	io.WriteString(w, svgEnd)
}
