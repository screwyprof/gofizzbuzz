package gofizbuzz_test

import (
	"testing"

	gofizbuzz "github.com/screwyprof/gofizzbuzz"
)

func TestFizzBuzzModuli(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		n    int
		want string
	}{
		{name: "It returns '1' given 1", n: 1, want: "1"},
		{name: "It returns '2' given 2", n: 2, want: "2"},
		{name: "It returns 'Fizz' given 3", n: 3, want: "Fizz"},
		{name: "It returns '4' given 4", n: 4, want: "4"},
		{name: "It returns 'Buzz' given 5", n: 5, want: "Buzz"},
		{name: "It returns 'FizzBuzz' given 15", n: 15, want: "FizzBuzz"},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := gofizbuzz.FizzBuzzModuli(tc.n)
			assertEquals(t, tc.want, got)
		})
	}
}

func BenchmarkFizzBuzzModuli(b *testing.B) {
	defer quiet()()

	for n := 0; n < b.N; n++ {
		gofizbuzz.FizzBuzzModuli(n)
	}
}
