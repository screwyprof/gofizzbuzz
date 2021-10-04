package gofizzbuzz

import "strconv"

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
	var res string

	for v, filter := range fizzBuzzFilters {
		if filter(n) {
			res += FizzBuzzValues[v]
		}
	}

	if res == "" {
		return strconv.Itoa(n)
	}

	return res
}
