package foldable

import "github.com/screwyprof/gofizzbuzz/monoid"

type (
	MonoidFoldable []monoid.OptionString
)

func (f MonoidFoldable) Foldl(init T, foldFunc func(result, next T) T) T {
	result := init
	for _, x := range f {
		result = foldFunc(result, x)
	}

	return result
}

func (f MonoidFoldable) Init() Foldable {
	return MonoidFoldable{}
}

func (f MonoidFoldable) Append(item T) Foldable {
	return append(f, item.(monoid.OptionString))
}
