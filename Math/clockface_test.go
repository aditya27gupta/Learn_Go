package clockface

import (
	"math"
	"testing"
	"time"
)

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(312, time.October, 28, hours, minutes, seconds, 0, time.UTC)
}

func roughlyEqualFloat64(a, b float64) bool {
	const threshold = 1e-7
	return math.Abs(a-b) < threshold
}

func roughlyEqualPoints(a, b Point) bool {
	return roughlyEqualFloat64(a.X, b.X) && roughlyEqualFloat64(a.Y, b.Y)
}

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		Time  time.Time
		angle float64
	}{
		{simpleTime(0, 0, 30), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{simpleTime(0, 0, 7), (math.Pi / 30) * 7},
	}

	for _, test := range cases {
		t.Run(test.Time.Format(time.TimeOnly), func(t *testing.T) {
			want := test.angle
			got := secondsInRadians(test.Time)
			if got != want {
				t.Fatalf("Got %v, Wanted %v", got, want)
			}
		})
	}
}

func TestSecondHandPoint(t *testing.T) {
	cases := []struct {
		Time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{X: 0, Y: -1}},
		{simpleTime(0, 0, 45), Point{X: -1, Y: 0}},
	}

	for _, test := range cases {
		t.Run(test.Time.Format(time.TimeOnly), func(t *testing.T) {
			got := secondHandPoint(test.Time)
			want := test.point
			if !roughlyEqualPoints(got, want) {
				t.Errorf("Got %v, Want %v", got, want)
			}
		})
	}
}

func TestSecondHand(t *testing.T) {
	cases := []struct {
		Time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 30), Point{X: 150, Y: 150 + 90}},
	}

	for _, test := range cases {
		t.Run(test.Time.Format(time.TimeOnly), func(t *testing.T) {
			got := SecondHand(test.Time)
			want := test.point
			if !roughlyEqualPoints(got, want) {
				t.Errorf("Got %v, Want %v", got, want)
			}
		})
	}
}

func TestMinutesInRadians(t *testing.T) {
	cases := []struct {
		Time  time.Time
		angle float64
	}{
		{simpleTime(0, 30, 0), math.Pi},
		{simpleTime(0, 0, 7), 7 * (math.Pi / (30 * 60))},
	}

	for _, test := range cases {
		t.Run(test.Time.Format(time.TimeOnly), func(t *testing.T) {
			want := test.angle
			got := minutesInRadians(test.Time)
			if got != want {
				t.Fatalf("Got %v, Wanted %v", got, want)
			}
		})
	}
}

func TestMinuteHandPoint(t *testing.T) {
	cases := []struct {
		Time  time.Time
		point Point
	}{
		{simpleTime(0, 30, 0), Point{X: 0, Y: -1}},
		{simpleTime(0, 45, 0), Point{X: -1, Y: 0}},
	}

	for _, test := range cases {
		t.Run(test.Time.Format(time.TimeOnly), func(t *testing.T) {
			got := minuteHandPoint(test.Time)
			want := test.point
			if !roughlyEqualPoints(got, want) {
				t.Errorf("Got %v, Want %v", got, want)
			}
		})
	}
}

func TestHoursInRadians(t *testing.T) {
	cases := []struct {
		Time  time.Time
		angle float64
	}{
		{simpleTime(6, 0, 0), math.Pi},
		{simpleTime(0, 0, 0), 0},
		{simpleTime(21, 0, 0), math.Pi * 1.5},
		{simpleTime(0, 1, 30), math.Pi / ((6 * 60 * 60) / 90)},
	}

	for _, test := range cases {
		t.Run(test.Time.Format(time.TimeOnly), func(t *testing.T) {
			want := test.angle
			got := hoursInRadians(test.Time)
			if !roughlyEqualFloat64(got, want) {
				t.Fatalf("Got %v, Wanted %v", got, want)
			}
		})
	}
}

func TestHourHandPoint(t *testing.T) {
	cases := []struct {
		Time  time.Time
		point Point
	}{
		{simpleTime(6, 0, 0), Point{X: 0, Y: -1}},
		{simpleTime(21, 0, 0), Point{X: -1, Y: 0}},
	}

	for _, test := range cases {
		t.Run(test.Time.Format(time.TimeOnly), func(t *testing.T) {
			got := hourHandPoint(test.Time)
			want := test.point
			if !roughlyEqualPoints(got, want) {
				t.Errorf("Got %v, Want %v", got, want)
			}
		})
	}
}
