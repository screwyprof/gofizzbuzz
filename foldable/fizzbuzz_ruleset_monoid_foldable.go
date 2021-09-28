package foldable

import "github.com/screwyprof/gofizzbuzz/monoid"

type (
	FizzBuzzRulesetMonoidFoldable []monoid.FizzBuzzRuleset
)

func (f FizzBuzzRulesetMonoidFoldable) Foldl(init T, foldFunc func(result, next T) T) T {
	result := init
	for _, x := range f {
		result = foldFunc(result, x)
	}

	return result
}

func (f FizzBuzzRulesetMonoidFoldable) Init() Foldable {
	return FizzBuzzRulesetMonoidFoldable{}
}

func (f FizzBuzzRulesetMonoidFoldable) Append(item T) Foldable {
	return append(f, item.(monoid.FizzBuzzRuleset))
}
