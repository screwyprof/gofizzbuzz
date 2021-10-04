package monoid_test

import (
	"testing"

	"github.com/genkami/dogs/types/option"

	"github.com/screwyprof/gofizzbuzz/fbtest"
	"github.com/screwyprof/gofizzbuzz/monoid"
)

func TestFizzBuzzRuleset(t *testing.T) {
	t.Parallel()

	// A set S equipped with a binary operation S × S → S,
	// which we will denote •, is a monoid if it satisfies the following two axioms:

	t.Run("it has valid identity", func(t *testing.T) {
		t.Parallel()
		n := 3

		fizz := func(n int) bool {
			return n%3 == 0
		}

		ruleMonoid := monoid.NewRuleSetMonoid()
		fizzRule := monoid.Rule(fizz, "Fizz")

		// There exists an element e in S such that for every element a in S,
		// the equations e • a = a and a • e = a hold.
		t.Logf("%s = %s",
			option.UnwrapOr[string](ruleMonoid.Combine(ruleMonoid.Empty(), fizzRule)(n), ""),
			option.UnwrapOr[string](ruleMonoid.Combine(fizzRule, ruleMonoid.Empty())(n), ""),
		)

		fbtest.Equals(t, fizzRule(n), ruleMonoid.Combine(ruleMonoid.Empty(), fizzRule)(n))
		fbtest.Equals(t,
			ruleMonoid.Combine(ruleMonoid.Empty(), fizzRule)(n),
			ruleMonoid.Combine(fizzRule, ruleMonoid.Empty())(n),
		)
	})

	t.Run("it has valid associativity", func(t *testing.T) {
		t.Parallel()
		n := 5

		fizz := func(n int) bool {
			return n%3 == 0
		}

		buzz := func(n int) bool {
			return n%5 == 0
		}

		bizz := func(n int) bool {
			return n%7 == 0
		}

		ruleMonoid := monoid.NewRuleSetMonoid()

		a := monoid.Rule(fizz, "Fizz")
		b := monoid.Rule(buzz, "Buzz")
		c := monoid.Rule(bizz, "Bizz")

		// For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.
		t.Logf("%s = %s",
			option.UnwrapOr[string](ruleMonoid.Combine(ruleMonoid.Combine(a, b), c)(n), ""),
			option.UnwrapOr[string](ruleMonoid.Combine(a, ruleMonoid.Combine(b, c))(n), ""),
		)

		fbtest.Equals(t,
			ruleMonoid.Combine(ruleMonoid.Combine(a, b), c)(n),
			ruleMonoid.Combine(a, ruleMonoid.Combine(b, c))(n),
		)
	})
}
