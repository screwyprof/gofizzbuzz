package gofizzbuzz_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz"
)

func TestFizzBuzzAcceptance(t *testing.T) {
	t.Parallel()

	fizzBuzzers := []func(n int) string{
		gofizzbuzz.FizzBuzz,
		gofizzbuzz.FizzBuzzFilter,
		gofizzbuzz.FizzBuzzModuli,
		gofizzbuzz.FizzBuzzFunctional,
	}

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

	for _, fizzBuzzer := range fizzBuzzers { //nolint:paralleltest
		for _, tc := range testCases {
			tc, fizzBuzzer := tc, fizzBuzzer

			t.Run(funcName(fizzBuzzer)+" "+tc.name, func(t *testing.T) {
				t.Parallel()

				got := fizzBuzzer(tc.n)
				assertEquals(t, tc.want, got)
			})
		}
	}
}