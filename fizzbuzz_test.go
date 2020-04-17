package gofizbuzz_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz"
)

// 1) Make it fail.
// 2) Make it pass.

func TestFizzBuzz(t *testing.T) {
	t.Run("It returns '1' given 1", func(t *testing.T) {
		got := gofizbuzz.FizzBuzz(1)
		if "1" != got {
			t.Fatalf("want '1', got %q", got)
		}
	})
}
