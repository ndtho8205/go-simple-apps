package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add 3 to 2", func(t *testing.T) {
		got := Add(2, 3)
		want := 5

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

}

func ExampleAdd() {
	fmt.Println(Add(1, 5))
	// Output: 6
}

func TestSubtract(t *testing.T) {
	t.Run("subtract 3 from 2", func(t *testing.T) {
		got := Subtract(2, 3)
		want := 2 - 3

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})

	t.Run("subtract 5 from 100", func(t *testing.T) {
		got := Subtract(100, 5)
		want := 100 - 5

		if got != want {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	})
}

func ExampleSubtract() {
	fmt.Println(Subtract(10, 5))
	// Output: 5
}
