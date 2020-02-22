package clockface

import (
	"math"
	"testing"
	"time"
)

func TestSecondHand(t *testing.T) {
	cases := []struct {
		time  time.Time
		point Point
	}{
		{simpleTime(0, 0, 0), Point{0, -1}},
		{simpleTime(0, 0, 30), Point{0, 1}},
		{simpleTime(0, 0, 15), Point{1, 0}},
		{simpleTime(0, 0, 45), Point{-1, 0}},
	}

	for _, test := range cases {
		got := SecondHand(test.time)
		want := test.point

		assertPointEqual(t, got, want)
	}
}

func TestSecondInRadiant(t *testing.T) {
	cases := []struct {
		time   time.Time
		radian float64
	}{
		{simpleTime(0, 0, 0), math.Pi},
		{simpleTime(0, 0, 30), 0},
		{simpleTime(0, 0, 45), -math.Pi / 2},
		{simpleTime(0, 0, 7), math.Pi * (23.0 / 30)},
	}

	for _, test := range cases {
		t.Run(test.time.Format("15:04:05"), func(t *testing.T) {
			got := secondInRadian(test.time)
			want := test.radian
			assertFloatEqual(t, got, want)
		})
	}
}

func simpleTime(hours, minutes, seconds int) time.Time {
	return time.Date(2020, time.February, 10, hours, minutes, seconds, 0, time.UTC)
}

func assertPointEqual(t *testing.T, got, want Point) {
	t.Helper()

	if !isEqual(got.X, want.X) || !isEqual(got.Y, want.Y) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertFloatEqual(t *testing.T, got, want float64) {
	t.Helper()

	if !isEqual(got, want) {
		t.Errorf("got %.4f want %.4f", got, want)
	}
}

func isEqual(got, want float64) bool {
	return math.Abs(got-want) < 1e-7
}
