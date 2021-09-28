package foldable

type (
	PredicateFoldable []func(n int) bool
	FoldFn            func(result, next T) T
)

func (f PredicateFoldable) Foldl(init T, foldFunc func(result, next T) T) T {
	result := init
	for _, x := range f {
		result = foldFunc(result, x)
	}

	return result
}

func (f PredicateFoldable) Init() Foldable {
	return PredicateFoldable{}
}

func (f PredicateFoldable) Append(item T) Foldable {
	return append(f, item.(func(n int) bool))
}
