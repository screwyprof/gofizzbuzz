package iterator

// Iterator iterates over some set of elements.
type Iterator[T any] interface {
	// Next returns the next element in this Iterable and advances its state.
	// The second return value is false if and only if there are no elements to return.
	Next() (T, bool)
}

// Fold accumulates every element in Iterator by applying fn.
func Fold[T, U any](init T, it Iterator[U], fn func(T, U) T) T {
	var acc = init
	for {
		x, ok := it.Next()
		if !ok {
			break
		}
		acc = fn(acc, x)
	}
	return acc
}

// Map returns an iterator that applies f to each element of it.
func Map[T, U any](it Iterator[T], fn func(T) U) Iterator[U] {
	return &mapIterator[T, U]{
		it: it,
		fn: fn,
	}
}

type mapIterator[T, U any] struct {
	it Iterator[T]
	fn func(T) U
}

func (it *mapIterator[T, U]) Next() (U, bool) {
	x, ok := it.it.Next()
	if !ok {
		var zero U
		return zero, false
	}
	return it.fn(x), true
}

// Slice is a slice with extra methods.
type Slice[T any] []T

// Iter returns an Iterator that iterates over given slice.
func (xs Slice[T]) Iter() Iterator[T] {
	return &sliceIterator[T]{
		xs:   ([]T)(xs),
		next: 0,
	}
}

type sliceIterator[T any] struct {
	xs   []T
	next int
}

func (it *sliceIterator[T]) Next() (T, bool) {
	if len(it.xs) <= it.next {
		var zero T
		return zero, false
	}
	i := it.next
	it.next++
	return it.xs[i], true
}
