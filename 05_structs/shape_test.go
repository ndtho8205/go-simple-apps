package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{3.5, 2.5}
		want := 12.0
		checkPerimeter(t, rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{1.0}
		want := 2 * math.Pi
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{2.0, 3.0}
		want := 6.0
		checkArea(t, rectangle, want)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		want := math.Pi * 100.0
		checkArea(t, circle, want)
	})
}

func TestAreaTDD(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{"Rectangle", Rectangle{3, 2}, 6.0},
		{"Circle", Circle{2.0}, math.Pi * 4.0},
		{"Square", Square{12}, 144.0},
	}

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got '%.2f' want '%.2f'", got, want)
		}
	}

	for _, test := range areaTests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.want)
		})
	}
}
