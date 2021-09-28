package monoid

type (
	FizzBuzzPredicate func(n int) bool
	FizzBuzzRuleset   func(n int) OptionString
)

func NewEmptyFizzBuzzRuleset() FizzBuzzRuleset {
	return func(n int) OptionString {
		return NoneString()
	}
}

func ForFizzBuzzPredicate(p FizzBuzzPredicate, word string) FizzBuzzRuleset {
	return func(n int) OptionString {
		if p(n) {
			return SomeString(word)
		}

		return NoneString()
	}
}

func (m FizzBuzzRuleset) Empty() FizzBuzzRuleset {
	return func(n int) OptionString {
		return NoneString()
	}
}

func (m FizzBuzzRuleset) Append(other FizzBuzzRuleset) FizzBuzzRuleset {
	return func(n int) OptionString {
		return m(n).Append(other(n))
	}
}
