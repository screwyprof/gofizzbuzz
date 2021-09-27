package gofizbuzz_test

import (
	"testing"

	gofizbuzz "github.com/screwyprof/gofizzbuzz"
)

func BenchmarkPrintFizzBuzzFunctional(b *testing.B) {
	defer quiet()()

	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		gofizbuzz.PrintFizzBuzzFuncional()
	}
}
