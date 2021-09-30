package monoid

import (
	"github.com/screwyprof/gofizzbuzz/iterator"
)

// Semigroup is a set of type `T` and its associative binary operation `Combine(T, T) T`
type Semigroup[T any] struct {
	Combine func(T, T) T
}

// SumWithInit sums up `init` and all values in `it`.
func (s Semigroup[T]) SumWithInit(init T, it iterator.Iterator[T]) T {
	return iterator.Fold[T, T](init, it, s.Combine)
}

//class Monoid a where
//mempty  :: a
//-- ^ Identity of 'mappend'
//mappend :: a -> a -> a
//-- ^ An associative operation
//mconcat :: [a] -> a

type Monoid[T any] struct {
	Item     T
	Identity func() T
	Semigroup[T]
}

func (m Monoid[T]) Append(a, b T) T {
	return m.Combine(a, b)
}

func (m Monoid[T]) Empty() T {
	return m.Identity()
}

func (m Monoid[T]) Unwrap() T {
	return m.Item
}
