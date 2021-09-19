package monoid

import "github.com/screwyprof/gofizzbuzz/option"

type OptionString option.String

func NewOptionString(s option.String) OptionString {
	return OptionString(s)
}

func NewOptionStringEmpty() OptionString {
	return OptionString(option.NoneString())
}

func (m OptionString) String() string {
	return option.String(m).UnwrapOrDefault()
}

func (m OptionString) Empty() OptionString {
	return OptionString(option.NoneString())
}

func (m OptionString) Append(other OptionString) OptionString {
	s := option.String(m).UnwrapOrDefault() + option.String(other).UnwrapOrDefault()

	return OptionString(option.NewString(&s))
}
