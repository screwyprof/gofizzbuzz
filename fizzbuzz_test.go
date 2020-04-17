package gofizbuzz_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz"
)

// 1) Make it fail. | RED
// 2) Make it pass. | GREEN
// 3) Make it good. | Refactor

func TestFizzBuzz(t *testing.T) {
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
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			got := gofizbuzz.FizzBuzz(tc.n)
			assertEquals(t, tc.want, got)
		})
	}
}

func assertEquals(t *testing.T, want, got string) {
	t.Helper()
	if want != got {
		t.Fatalf("want %q, got %q", want, got)
	}
}
