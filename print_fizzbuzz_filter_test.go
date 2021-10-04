package gofizzbuzz_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz"
	"github.com/screwyprof/gofizzbuzz/fbtest"
)

func BenchmarkPrintFizzBuzzFilter(b *testing.B) {
	defer fbtest.Quiet()()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		gofizzbuzz.PrintFizzBuzzFilter()
	}
}