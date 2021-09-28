package gofizzbuzz_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz"
)

func BenchmarkFizzBuzz(b *testing.B) {
	fizzBuzzers := []func(n int) string{
		gofizzbuzz.FizzBuzz,
		gofizzbuzz.FizzBuzzFilter,
		gofizzbuzz.FizzBuzzModuli,
		gofizzbuzz.FizzBuzzFunctional,
	}

	for _, fizzBuzzer := range fizzBuzzers {
		b.Run(funcName(fizzBuzzer), func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				fizzBuzzer(n)
			}
		})
	}
}
