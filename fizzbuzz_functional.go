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
	m := monoid.ForString("").
		Append(monoid.ForString("Fizz").Filtered(func() bool {
			return n%3 == 0
		})).
		Append(monoid.ForString("Buzz").Filtered(func() bool {
			return n%5 == 0
		}))

	return m.UnwrapOr(strconv.Itoa(n))

	//fizz := func(in interface{}) bool {
	//	return in.(int)%3 == 0
	//}
	//
	//buzz := func(in interface{}) bool {
	//	return in.(int)%5 == 0
	//}

	//predicates := []monoid.Predicate {
	//	monoid.ForPredicate(fizz),
	//	monoid.ForPredicate(buzz),
	//}
	//
	//// fold over predicates:  mconcat = foldr mappend mempty
	//composedPredicates := monoid.ForPredicate(nil)
	//for _, p := range predicates {
	//	p.Append(p)
	//}
	//
	////if composedPredicates(n).) {
	////
	////}
}

// 1. Generate a list
// 		list = [1, 2, 3, 4, 5]
// 2. Somehow compose func(int) string as a Monoid?
//    mconcat :: [a] -> a
//    mconcat = foldr mappend mempty
// 3. FizzBuzzMonoid.Mconcat()
