package gofizbuzz

import (
	"strconv"

	"github.com/screwyprof/gofizzbuzz/monoid"
)

// FizzBuzzFunctional
//
// https://web.archive.org/web/20130511210903/http://dave.fayr.am/posts/2012-10-4-finding-fizzbuzz.html
// https://medium.com/@iopguy/fizzbuzz-can-finally-be-implemented-in-stable-rust-87649a882f2d
func FizzBuzzFunctional(n int) string {
	fizzFilter := func() bool {
		return n%3 == 0
	}

	buzzFilter := func() bool {
		return n%5 == 0
	}

	m := monoid.NewEmptyString().
		AppendIf(fizzFilter, "Fizz").
		AppendIf(buzzFilter, "Buzz")

	return FromMaybe(n, m)
}

func FromMaybe(n int, m monoid.String) string {
	if m == m.Empty() {
		return strconv.Itoa(n)
	}

	return m.String()
}
