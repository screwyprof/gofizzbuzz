package gofizzbuzz

import (
	"strconv"

	"github.com/screwyprof/gofizzbuzz/fastdiv"
)

type fizzBuzzFilter = func(n int) bool

var fizzFilter = func(n int) bool {
	return fastdiv.IsDivisible(uint64(n), fastdiv.M3)
}

var buzzFilter = func(n int) bool {
	return fastdiv.IsDivisible(uint64(n), fastdiv.M5)
}

var fizzBuzzFilters = []fizzBuzzFilter{
	fizzFilter,
	buzzFilter,
}

var FizzBuzzValues = []string{"Fizz", "Buzz"}

func FizzBuzzFilter(n int) string {
	var res string

	for v, filter := range fizzBuzzFilters {
		if filter(n) {
			res += FizzBuzzValues[v]
		}
	}

	if res != "" {
		return res
	}

	return strconv.FormatInt(int64(n), 10)
}
