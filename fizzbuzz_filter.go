package gofizzbuzz

import (
	"strconv"

	"github.com/screwyprof/gofizzbuzz/fastdiv"
)

type fizzBuzzPredicate = func(n int) bool

var fizzBuzzFilter = func(n int) bool {
	return fastdiv.IsDivisible(uint64(n), fastdiv.M15)
}

var fizzFilter = func(n int) bool {
	return fastdiv.IsDivisible(uint64(n), fastdiv.M3)
}

var buzzFilter = func(n int) bool {
	return fastdiv.IsDivisible(uint64(n), fastdiv.M5)
}

var fizzBuzzFilters = []fizzBuzzPredicate{
	fizzBuzzFilter,
	fizzFilter,
	buzzFilter,
}

var FizzBuzzValues = []string{"FizzBuzz", "Fizz", "Buzz"}

func FizzBuzzFilter(n int) string {
	for v, filter := range fizzBuzzFilters {
		if filter(n) {
			return FizzBuzzValues[v]
		}
	}

	return strconv.FormatInt(int64(n), 10)
}
