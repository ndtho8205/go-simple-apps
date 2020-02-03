package slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{-1, 1, 0, 1, -1, 2}

	got := Sum(numbers)
	want := 2

	if got != want {
		t.Errorf("given %v got '%d' want '%d'", numbers, got, want)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3}, []int{-1, 1, 2})
	want := []int{6, 2}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {
	assertSumsEqual := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum the tails of non-empty slices", func(t *testing.T) {
		got := SumAllTails([]int{1}, []int{1, 2}, []int{1, 2, 3})
		want := []int{0, 2, 5}
		assertSumsEqual(t, got, want)
	})

	t.Run("sum the tails of an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{})
		want := []int{0}
		assertSumsEqual(t, got, want)
	})
}
