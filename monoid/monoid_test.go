package monoid_test

import (
	"github.com/genkami/dogs/classes/algebra"
	"github.com/genkami/dogs/types/iterator"
	"github.com/genkami/dogs/types/option"
	"github.com/genkami/dogs/types/slice"
	"github.com/screwyprof/gofizzbuzz/monoid/monoidtest"
	"strconv"
	"testing"
)

//
//import (
//	"github.com/screwyprof/gofizzbuzz/iterator"
//	"github.com/screwyprof/gofizzbuzz/monoid"
//	"github.com/screwyprof/gofizzbuzz/monoid/monoidtest"
//	"github.com/screwyprof/gofizzbuzz/option"
//	"testing"
//)
//
//func TestToStringMonoid(t *testing.T) {
//	t.Parallel()
//
//	// A set S equipped with a binary operation S × S → S,
//	// which we will denote •, is a monoid if it satisfies the following two axioms:
//	t.Run("it has valid identity", func(t *testing.T) {
//		t.Parallel()
//
//		m := ToStringMonoid("TEST")
//
//		// There exists an element e in S such that for every element a in S,
//		// the equations e • a = a and a • e = a hold.
//
//		monoidtest.AssertEqual(t, m.Item, m.Append(m.Empty(), m.Item))
//		monoidtest.AssertEqual(t, m.Append(m.Empty(), m.Item), m.Append(m.Item, m.Empty()))
//	})
//
//	t.Run("it has valid associativity", func(t *testing.T) {
//		t.Parallel()
//
//		m := ToStringMonoid("")
//
//		a := "a"
//		b := "b"
//		c := "c"
//
//		// For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.
//		monoidtest.AssertEqual(t, m.Append(m.Append(a, b), c), m.Append(a, m.Append(b, c)))
//	})
//}
//
//func TestToOptionalStringMonoid(t *testing.T) {
//	t.Parallel()
//
//	// A set S equipped with a binary operation S × S → S,
//	// which we will denote •, is a monoid if it satisfies the following two axioms:
//	t.Run("it has valid identity", func(t *testing.T) {
//		t.Parallel()
//
//		m := ToOptionStringMonoid(option.SomeString("TEST"))
//
//		// There exists an element e in S such that for every element a in S,
//		// the equations e • a = a and a • e = a hold.
//
//		monoidtest.AssertEqual(t, m.Item, m.Append(m.Empty(), m.Item))
//		monoidtest.AssertEqual(t, m.Append(m.Empty(), m.Item), m.Append(m.Item, m.Empty()))
//	})
//
//	t.Run("it has valid associativity", func(t *testing.T) {
//		t.Parallel()
//
//		m := ToOptionStringMonoid(option.NoneString())
//
//		a := option.SomeString("foo")
//		b := option.NoneString()
//		c := option.SomeString("bar")
//
//		// For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.
//		monoidtest.AssertEqual(t, m.Append(m.Append(a, b), c), m.Append(a, m.Append(b, c)))
//	})
//}
//
//func TestFizzBuzzRuleSetMonoid(t *testing.T) {
//	t.Parallel()
//
//	// A set S equipped with a binary operation S × S → S,
//	// which we will denote •, is a monoid if it satisfies the following two axioms:
//	t.Run("it has valid identity", func(t *testing.T) {
//		t.Parallel()
//
//		n := 3
//
//		fizz := func(n int) bool {
//			return n%3 == 0
//		}
//
//		buzz := func(n int) bool {
//			return n%5 == 0
//		}
//
//		m := ToRuleSet(
//			Rule(fizz, "Fizz"),
//			Rule(buzz, "Buzz"),
//		)
//
//		// There exists an element e in S such that for every element a in S,
//		// the equations e • a = a and a • e = a hold.
//		monoidtest.AssertEqual(t, m.Append(m.Item, m.Empty())(n), m.Item(n))
//		monoidtest.AssertEqual(t, m.Append(m.Empty(), m.Item)(n), m.Append(m.Item, m.Empty())(n))
//		//monoidtest.AssertEqual(t, m(n).Append(m(n).Empty(), m(n).Item), m(n).Append(m(n).Item, m(n).Empty()))
//	})
//	//
//	//t.Run("it has valid associativity", func(t *testing.T) {
//	//	t.Parallel()
//	//	n := 3
//	//
//	//	fizz := func(n int) bool {
//	//		return n%3 == 0
//	//	}
//	//
//	//	buzz := func(n int) bool {
//	//		return n%5 == 0
//	//	}
//	//
//	//	bizz := func(n int) bool {
//	//		return n%7 == 0
//	//	}
//	//
//	//	a := ToFizzBuzzRuleSetMonoid(fizz, "Fizz")
//	//	b := ToFizzBuzzRuleSetMonoid(buzz, "Buzz")
//	//	c := ToFizzBuzzRuleSetMonoid(bizz, "Bizz")
//	//
//	//	// For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.
//	//	monoidtest.AssertEqual(t, a.Append(b).Append(c)(n), a.Append(b.Append(c))(n))
//	//})
//}
//
//func ToStringMonoid(s string) monoid.Monoid[string] {
//	return monoid.Monoid[string]{
//		Semigroup: monoid.Semigroup[string]{
//			Combine: func(a, b string) string { return a + b },
//		},
//		Item:     s,
//		Identity: func() string { return "" },
//	}
//}
//
//func ToOptionStringMonoid(o option.String) monoid.Monoid[option.String] {
//	return monoid.Monoid[option.String]{
//		Item:     o,
//		Identity: func() option.String { return option.NoneString() },
//		Semigroup: monoid.Semigroup[option.String]{
//			Combine: func(a, b option.String) option.String {
//				switch {
//				case a == option.NoneString():
//					return b
//				case b == option.NoneString():
//					return a
//				default:
//					return option.SomeString(a.UnwrapOrDefault() + b.UnwrapOrDefault())
//				}
//			},
//		},
//	}
//}
//
//type (
//	FizzBuzzPredicate func(n int) bool
//	//FizzBuzzRuleset monoid.Monoid[func(n int) monoid.Monoid[option.String]]
//	//FizzBuzzRuleset monoid.Monoid[[]FizzBuzzRulesetFn]
//
//	FizzBuzzRule = func(n int) option.String
//)
//
//func Rule(p FizzBuzzPredicate, word string) FizzBuzzRule {
//	return func(n int) option.String {
//		if p(n) {
//			return option.SomeString(word)
//		}
//
//		return option.NoneString()
//	}
//}
//
//func ToRuleMonoid(rule FizzBuzzRule) monoid.Monoid[FizzBuzzRule] {
//	m := monoid.Monoid[FizzBuzzRule]{
//		Item: rule,
//		Identity: func() FizzBuzzRule {
//			return func(n int) option.String {
//				return option.NoneString()
//			}
//		},
//		Semigroup: monoid.Semigroup[FizzBuzzRule]{
//			Combine: func(a, b FizzBuzzRule) FizzBuzzRule {
//				switch {
//				case a == nil:
//					return b
//				case b == nil:
//					return a
//				default:
//					return func(n int) option.String {
//						return a(n).Append(b(n))
//					}
//				}
//			},
//		},
//	}
//
//	return m
//}
//
////
////type sliceIterator[T any] struct {
////	xs []T
////	next int
////}
////
////
////func (it *sliceIterator[T]) Next() (T, bool) {
////	if len(it.xs) <= it.next {
////		var zero T
////		return zero, false
////	}
////	i := it.next
////	it.next++
////	return it.xs[i], true
////}
//
////func ToRulesSetIterator[T any](rules []FizzBuzzRule) iterator.Iterator[T]{
////	return &RuleSetIterator[T]{
////		xs: rules,
////	}
////}
////
////type RuleSetIterator[T any] struct {
////	xs []T
////	next int
////}
////
////func (it *RuleSetIterator[T]) Next() (monoid.Monoid[T], bool) {
////	if len(it.xs) <= it.next {
////		return ToRuleMonoid(nil), false
////	}
////
////	fn := func (from FizzBuzzRule) monoid.Monoid[T] {
////		return ToRuleMonoid(from)
////	}
////
////	i := it.next
////	it.next++
////
////	return fn(it.xs[i]), true
////}
//
//func toRulesSetIterator(xs ...FizzBuzzRule) iterator.Iterator[FizzBuzzRule] {
//	return iterator.Slice[FizzBuzzRule](xs).Iter()
//}
//
//func ToRuleSet(rule ...FizzBuzzRule) monoid.Monoid[FizzBuzzRule] {
//	it := toRulesSetIterator(rule...)
//
//
//	fn := func(acc monoid.Monoid[FizzBuzzRule], next FizzBuzzRule) monoid.Monoid[FizzBuzzRule] {
//		//return acc.Append(acc.Empty(), ToRuleMonoid(next).Item)
//		//fmt.Printf("type: %v\n", reflect.TypeOf(next).String())
//		//fmt.Printf("%+#v\n", next(3).UnwrapOrDefault())
//		//fmt.Printf("%+#v\n", next(5).UnwrapOrDefault())
//
//		res := acc.Append(acc.Empty(), ToRuleMonoid(next).Item)
//		return ToRuleMonoid(res)
//	}
//
//	return iterator.Fold(ToRuleMonoid(nil), it, fn)
//
//	//return m.SumWithInit(ToRuleMonoid(nil), it)
//}
//
////  fizzbuzz i = fromMaybe (show i) (ruleSet i)
////	  where
////		 ruleSet = fold filters

func TestFizzBuzz(t *testing.T) {
	t.Parallel()
	fizz := func(n int) bool {
		return n%3 == 0
	}

	buzz := func(n int) bool {
		return n%5 == 0
	}

	filters := []FizzBuzzRuleFn{
		Rule(fizz, "Fizz"),
		Rule(buzz, "Buzz"),
	}

	ruleSet := RuleSet(filters...)

	monoidtest.AssertEqual(t, "Fizz", option.UnwrapOr(ruleSet(3), strconv.Itoa(3)))
	monoidtest.AssertEqual(t, "Buzz", option.UnwrapOr(ruleSet(5), strconv.Itoa(5)))
	monoidtest.AssertEqual(t, "7", option.UnwrapOr(ruleSet(7), strconv.Itoa(7)))
}

type FizzBuzzRuleFn func(n int) option.Option[string]

func Rule(p func(n int) bool, word string) FizzBuzzRuleFn {
	return func(n int) option.Option[string] {
		if p(n) {
			return option.Some(word)
		}

		return option.None[string]()
	}
}

type FizzBuzzRuleSet interface {
	~func(n int) option.Option[string]
}

// DeriveFizzBuzzRuleSetSemigroup derives Semigroup using `+` operator.
func DeriveFizzBuzzRuleSetSemigroup[T FizzBuzzRuleSet]() algebra.Semigroup[T] {
	return fizzBuzzRuleSetSemigroup[T]{}
}

type fizzBuzzRuleSetSemigroup[T FizzBuzzRuleSet] struct{}

func (fizzBuzzRuleSetSemigroup[T]) Combine(x, y T) T {
	return func(n int) option.Option[string] {
		xSome := x(n)
		ySome := y(n)

		switch {
		case !option.IsSome(xSome):
			return ySome
		case !option.IsSome(ySome):
			return xSome
		default:
			return option.Some(option.UnwrapOr(xSome, "") + option.UnwrapOr(ySome, ""))
		}
	}
}

func DeriveMonoid[T any](s algebra.Semigroup[FizzBuzzRuleFn]) algebra.Monoid[FizzBuzzRuleFn] {
	return &algebra.DefaultMonoid[FizzBuzzRuleFn]{
		Semigroup: s,
		EmptyImpl: func() FizzBuzzRuleFn {
			return func(n int) option.Option[string] {
				return option.None[string]()
			}
		},
	}
}

func RuleSet(rule ...FizzBuzzRuleFn) FizzBuzzRuleFn {
	m := DeriveMonoid[FizzBuzzRuleFn](DeriveFizzBuzzRuleSetSemigroup[FizzBuzzRuleFn]())

	it := slice.Slice[FizzBuzzRuleFn](rule).Iter()
	return iterator.Sum[FizzBuzzRuleFn](it, m)
}
