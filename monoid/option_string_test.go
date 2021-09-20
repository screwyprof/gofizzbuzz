package monoid_test

import (
	"testing"

	"github.com/screwyprof/gofizzbuzz/monoid"
	"github.com/screwyprof/gofizzbuzz/monoid/monoidtest"
)

func TestOptionString(t *testing.T) {
	t.Parallel()

	// A set S equipped with a binary operation S × S → S,
	// which we will denote •, is a monoid if it satisfies the following two axioms:

	t.Run("it has valid identity", func(t *testing.T) {
		t.Parallel()

		m := monoid.SomeString("test")

		// There exists an element e in S such that for every element a in S,
		// the equations e • a = a and a • e = a hold.
		monoidtest.AssertEqual(t, m, m.Append(m.Empty()))
		monoidtest.AssertEqual(t, m.Empty().Append(m), m.Append(m.Empty()))
	})

	t.Run("it has valid associativity", func(t *testing.T) {
		t.Parallel()

		a := monoid.SomeString("foo")
		b := monoid.SomeString("bar")
		c := monoid.NoneString()

		// For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.
		monoidtest.AssertEqual(t, a.Append(b).Append(c), a.Append(b.Append(c)))
	})
}
