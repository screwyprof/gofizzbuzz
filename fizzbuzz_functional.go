package gofizzbuzz

import (
	"strconv"

	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/types/option"
	"github.com/genkami/dogs/types/slice"
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
	fizz := func(n int) bool {
		return n%3 == 0
	}

	buzz := func(n int) bool {
		return n%5 == 0
	}

	rules := slice.Slice[FizzBuzzRule]{
		Rule(fizz, "Fizz"),
		Rule(buzz, "Buzz"),
	}

	ruleMonoid := NewRuleSetMonoid()
	slice.Sum(rules, ruleMonoid)

	ruleSet := slice.Sum(rules, ruleMonoid)

	return option.UnwrapOr(ruleSet(n), strconv.FormatInt(int64(n), 10))
}

type FizzBuzzRule func(n int) option.Option[string]

func Rule(p func(n int) bool, word string) FizzBuzzRule {
	return func(n int) option.Option[string] {
		if p(n) {
			return option.Some(word)
		}

		return option.None[string]()
	}
}

func NewRuleSetMonoid() algebra.Monoid[FizzBuzzRule] {
	optionMonoid := option.DeriveMonoid[string](algebra.DeriveAdditiveSemigroup[string]())

	ruleMonoid := &algebra.DefaultMonoid[FizzBuzzRule]{
		Semigroup: &algebra.DefaultSemigroup[FizzBuzzRule]{
			CombineImpl: func(a, b FizzBuzzRule) FizzBuzzRule {
				return func(n int) option.Option[string] {
					return optionMonoid.Combine(a(n), b(n))
				}
			},
		},
		EmptyImpl: func() FizzBuzzRule {
			return func(_ int) option.Option[string] {
				return option.None[string]()
			}
		},
	}

	return ruleMonoid
}
