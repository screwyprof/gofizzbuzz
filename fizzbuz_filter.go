package gofizbuzz

import (
	"strconv"

	"github.com/screwyprof/gofizzbuzz/monoid"
)

type fizzBuzzFilter = func(n int) bool

var fizzFilter = func(n int) bool {
	return n%3 == 0
}

var buzzFilter = func(n int) bool {
	return n%5 == 0
}

var fizzBuzzFilters = []fizzBuzzFilter{
	fizzFilter,
	buzzFilter,
}

var FizzBuzzValues = []string{"Fizz", "Buzz"}

func FizzBuzzFilter(n int) string {
	res := monoid.NewString("")

	for v, filter := range fizzBuzzFilters {
		if filter(n) {
			res = res.Append(monoid.NewString(FizzBuzzValues[v]))
		}
	}

	if res == res.Empty() {
		return strconv.Itoa(n)
	}

	return res.String()
}
