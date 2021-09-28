package gofizzbuzz

import (
	"strconv"

	"github.com/screwyprof/gofizzbuzz/foldable"
	"github.com/screwyprof/gofizzbuzz/monoid"
)

// FizzBuzzFunctional
//
// https://web.archive.org/web/20130511210903/http://dave.fayr.am/posts/2012-10-4-finding-fizzbuzz.html
// https://medium.com/@iopguy/fizzbuzz-can-finally-be-implemented-in-stable-rust-87649a882f2d
// https://www.parsonsmatt.org/2016/02/27/an_elegant_fizzbuzz.html
// https://blog.ploeh.dk/2017/10/10/strings-lists-and-sequences-as-a-monoid/
// https://blog.ploeh.dk/2017/11/06/function-monoids/
// https://blog.ploeh.dk/2017/10/06/monoids/
//
//  fizzbuzz i = fromMaybe (show i) (ruleSet i)
//	  where
//		 ruleSet = fold filters
func FizzBuzzFunctional(n int) string {
	// fizzbuzz i = fromMaybe (show i) $ mappend ["fizz" | i `rem` 3 == 0]
	//											["buzz" | i `rem` 5 == 0]

	//m := monoid.ForString("").
	//	Append(monoid.ForString("Fizz").Filtered(func() bool {
	//		return n%3 == 0
	//	})).
	//	Append(monoid.ForString("Buzz").Filtered(func() bool {
	//		return n%5 == 0
	//	}))
	//
	//return m.UnwrapOr(strconv.Itoa(n))

	fizz := func(n int) bool {
		return n%3 == 0
	}

	buzz := func(n int) bool {
		return n%5 == 0
	}

	filters := foldable.RuleFoldable{
		monoid.ForFizzBuzzPredicate(fizz, "Fizz"),
		monoid.ForFizzBuzzPredicate(buzz, "Buzz"),
	}

	rules := fold(filters, monoid.NewEmptyFizzBuzzRuleset())

	return fromOption(strconv.Itoa(n), rules(n))
}

// fold :: (Foldable t, Monoid m) => t m -> m
func fold(filters foldable.RuleFoldable, m monoid.FizzBuzzRuleset) monoid.FizzBuzzRuleset {
	rules := filters.Foldl(filters.Init(), func(result foldable.T, next foldable.T) foldable.T {
		rule, ok := next.(func(int) monoid.OptionString)
		mustOk(ok, "cannot cast result foldable.T to func(int) monoid.OptionString")

		m = m.Append(rule)

		return m
	})

	ruleSet, ok := rules.(monoid.FizzBuzzRuleset)
	mustOk(ok, "cannot cast result foldable.T to monoid.FizzBuzzRuleset")

	return ruleSet
}

func fromOption(s string, o monoid.OptionString) string {
	return o.UnwrapOr(s)
}

func mustOk(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}
