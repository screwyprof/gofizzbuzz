package gofizbuzz

import (
	"github.com/screwyprof/gofizzbuzz/monoid"
	"strconv"
)

// FizzBuzzFunctional
//
// https://web.archive.org/web/20130511210903/http://dave.fayr.am/posts/2012-10-4-finding-fizzbuzz.html
// https://medium.com/@iopguy/fizzbuzz-can-finally-be-implemented-in-stable-rust-87649a882f2d
func FizzBuzzFunctional(n int) string {
	//m := monoid.NoneString().
	//	Append(monoid.SomeString("Fizz").Filtered(func() bool {
	//		return n%3 == 0
	//	})).
	//	Append(monoid.SomeString("Buzz").Filtered(func() bool {
	//		return n%5 == 0
	//	}))

	//return m.UnwrapOr(strconv.Itoa(n))

	m := monoid.NewString("").
		Append(monoid.NewString("Fizz").Filtered(func() bool {
			return n%3 == 0
		})).
		Append(monoid.NewString("Buzz").Filtered(func() bool {
			return n%5 == 0
		}))

	return m.UnwrapOr(strconv.Itoa(n))
}
