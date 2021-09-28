package gofizzbuzz_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz"
)

func BenchmarkPrintFizzBuzz(b *testing.B) {
	defer quiet()()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		gofizzbuzz.PrintFizzBuzz()
	}
}
