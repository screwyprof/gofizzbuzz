package monoid

import (
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/types/option"
)

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
