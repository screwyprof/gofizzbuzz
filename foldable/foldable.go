package foldable

type T interface{}

type Foldable interface {
	Foldl(init T, f func(result, next T) T) T
	Init() Foldable
	Append(item T) Foldable
}

// Concat concatenates the parameters.
func Concat(a, b Foldable) Foldable {
	return b.Foldl(a, func(result, next T) T {
		return result.(Foldable).Append(next)
	}).(Foldable)
}

// Map applies a function to each item inside the Foldable.
func Map(f Foldable, mapFunc func(T) T) Foldable {
	return f.Foldl(f.Init(), func(result, next T) T {
		return result.(Foldable).Append(mapFunc(next))
	}).(Foldable)
}
