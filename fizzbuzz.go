package gofizzbuzz

import (
	"strconv"

	"github.com/screwyprof/gofizzbuzz/fastdiv"
)

func FizzBuzz(n int) string {
	switch {
	case fastdiv.IsDivisible(uint64(n), fastdiv.M15):
		return "FizzBuzz"
	case fastdiv.IsDivisible(uint64(n), fastdiv.M5):
		return "Buzz"
	case fastdiv.IsDivisible(uint64(n), fastdiv.M3):
		return "Fizz"
	default:
		return strconv.FormatInt(int64(n), 10)
	}
}
