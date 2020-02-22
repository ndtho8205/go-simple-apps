package clockface

import (
	"math"
	"time"
)

const (
	ClockRadius      = 0
	SecondHandLength = 1
)

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	angle := secondInRadian(t)
	point := Point{
		X: math.Sin(angle)*SecondHandLength + ClockRadius,
		Y: math.Cos(angle)*SecondHandLength + ClockRadius,
	}

	return point
}

func secondInRadian(t time.Time) float64 {
	second := float64(t.Second())
	return math.Pi * (1 - second/30)
}
