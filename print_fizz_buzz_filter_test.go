package gofizzbuzz_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz"
)

func BenchmarkPrintFizzBuzzFilter(b *testing.B) {
	defer quiet()()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		gofizzbuzz.PrintFizzBuzzFilter()
	}
}
