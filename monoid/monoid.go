package monoid

// Semigroup is a set of type `T` and its associative binary operation `Combine(T, T) T`
//type Semigroup[T any] interface {
//	Append(T, T) T
//}

//// Monoid is a Semigroup with identity.
//type Monoid[T any] interface {
//	Empty() T
//	Append(other T) T
//}

//class Monoid a where
//mempty  :: a
//-- ^ Identity of 'mappend'
//mappend :: a -> a -> a
//-- ^ An associative operation
//mconcat :: [a] -> a

//type StringMonoidBuilder[T any] struct {
//	Identity func() T
//	Combine func (T, T) T
//}
//
//func (b StringMonoidBuilder[T]) Build() GenericMonoid[T] {
//	return GenericMonoid[T]{
//		identity: b.Identity(),
//		combine: b.Combine,
//	}
//}
//
//type GenericMonoid[T any] struct {
//	identity T
//	combine func (T, T) T
//}
//
//func (m GenericMonoid[T]) Empty() Monoid[T] {
//	return GenericMonoid[T]{
//		identity: m.identity,
//		combine: m.combine,
//	}
//}
//
//func (m GenericMonoid[T]) Append(other Monoid[T]) Monoid[T] {
//	 //self := GenericMonoid[T]{
//		//identity: m.identity,
//		//combine: m.combine,
//	 //}
//	 //
//	 //self.
//	//return m.combine(m, other)
//}

//func (m GenericMonoid[T]) Unwrap() T {
//	return m.Item
//}

//type GenericMonoid[T any] struct{
//	empty func() T
//}
//
//func (m GenericMonoid[T]) Empty() T {
//	return m.empty()
//}
//
//func (m GenericMonoid[T]) Append(other T) T {
//	return m.item + other
//}

//func ForNewString(s string) Monoid[GenericStringMonoid[string]] {
//	return GenericStringMonoid[string]{item: s}
//}

//type PredicatFn[Cond any] func(Cond) bool
//
//type Predicate[Cond any] interface {
//	func(Cond) bool
//}
//
//// DerivePredicateSemigroup derives Semigroup using `+` operator.
//func DerivePredicateSemigroup[Cond any, T Predicate[Cond]]() algebra.Semigroup[T] {
//	return predicateSemigroup[Cond,T]{}
//}
//
//type predicateSemigroup[Cond any, T Predicate[Cond]] struct{}
//
//func (predicateSemigroup[Cond,T]) Combine(x, y T) T {
//	return func(c Cond) bool {
//		fmt.Printf("c: %v\n", c)
//		fmt.Printf("x(c): %v\n", x(c))
//		fmt.Printf("y(c): %v\n", y(c))
//
//		return x(c) && y(c)
//	}
//}
