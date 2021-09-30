package monoid_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz/monoid"
	"github.com/screwyprof/gofizzbuzz/monoid/monoidtest"
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

		m := monoid.ForFizzBuzzPredicate(fizz, "Fizz")

		// There exists an element e in S such that for every element a in S,
		// the equations e • a = a and a • e = a hold.

		// t.Logf("%v=%v", m(n).UnwrapOr(""), m.Append(m.Empty())(n).UnwrapOr(""))
		// t.Logf("%v=%v", m.Empty().Append(m)(n).UnwrapOr(""), m.Append(m.Empty())(n).UnwrapOr(""))
		monoidtest.AssertEqual(t, m(n), m.Append(m.Empty())(n))
		monoidtest.AssertEqual(t, m.Empty().Append(m)(n), m.Append(m.Empty())(n))
	})

	t.Run("it has valid associativity", func(t *testing.T) {
		t.Parallel()
		n := 3

		fizz := func(n int) bool {
			return n%3 == 0
		}

		buzz := func(n int) bool {
			return n%5 == 0
		}

		bizz := func(n int) bool {
			return n%7 == 0
		}

		a := monoid.ForFizzBuzzPredicate(fizz, "Fizz")
		b := monoid.ForFizzBuzzPredicate(buzz, "Buzz")
		c := monoid.ForFizzBuzzPredicate(bizz, "Bizz")

		a.Append(b)(3)

		// For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.

		// t.Logf("%v=%v", a.Append(b).Append(c)(n).UnwrapOr(""), a.Append(b.Append(c))(n).UnwrapOr(""))
		monoidtest.AssertEqual(t, a.Append(b).Append(c)(n), a.Append(b.Append(c))(n))
	})
}
