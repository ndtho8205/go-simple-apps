package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	assertEqual := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("repeat an empty string 2 times", func(t *testing.T) {
		got := Repeat("", 2)
		want := ""
		assertEqual(t, got, want)
	})

	t.Run("repeat the character 'a' 0 times", func(t *testing.T) {
		got := Repeat("a", 0)
		want := ""
		assertEqual(t, got, want)
	})

	t.Run("repeat the character 'a' 1 times", func(t *testing.T) {
		got := Repeat("a", 1)
		want := "a"
		assertEqual(t, got, want)
	})

	t.Run("repeat the character 'a' 5 times", func(t *testing.T) {
		got := Repeat("a", 5)
		want := "aaaaa"
		assertEqual(t, got, want)
	})

	t.Run("repeat the string 'hello' 7 times", func(t *testing.T) {
		got := Repeat("hello", 7)
		want := "hellohellohellohellohellohellohello"
		assertEqual(t, got, want)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("hello", 5)
	}
}

func BenchmarkStandardRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StandardRepeat("hello", 5)
	}
}

func ExampleRepeat() {
	fmt.Println(Repeat("a", 3))
	// Output: aaa
}
