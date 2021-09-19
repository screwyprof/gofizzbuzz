package gofizbuzz

import (
	"strconv"
	"strings"

	"github.com/screwyprof/gofizzbuzz/monoid"
	"github.com/screwyprof/gofizzbuzz/option"
)

func FizzBuzzFunctional(n int) string {
	fizzFilter := func(n int) func() bool {
		return func() bool {
			return n%3 == 0
		}
	}

	buzzFilter := func(n int) func() bool {
		return func() bool {
			return n%5 == 0
		}
	}

	fizzOp := option.NewStringIf(fizzFilter(n), "Fizz")
	buzzOp := option.NewStringIf(buzzFilter(n), "Buzz")

	m := OptionToStringMonoid(fizzOp).
		Append(OptionToStringMonoid(buzzOp))

	if m.String() == "" {
		return strconv.Itoa(n)
	}

	return strings.TrimSpace(m.String())
}

func OptionToStringMonoid(op option.String) monoid.OptionString {
	return monoid.NewOptionString(op)
}
