package foldable

import "github.com/screwyprof/gofizzbuzz/monoid"

type (
	RuleFoldable []func(n int) monoid.OptionString
)

func (f RuleFoldable) Foldl(init T, foldFunc func(result, next T) T) T {
	result := init
	for _, x := range f {
		result = foldFunc(result, x)
	}

	return result
}

func (f RuleFoldable) Init() Foldable {
	return RuleFoldable{}
}

func (f RuleFoldable) Append(item T) Foldable {
	return append(f, item.(func(n int) monoid.OptionString))
}
