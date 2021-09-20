package monoid

import "github.com/screwyprof/gofizzbuzz/option"

type OptionString struct {
	op option.String
}

// SomeString wraps the s into a String.
func SomeString(s string) OptionString {
	return OptionString{
		op: option.SomeString(s),
	}
}

// NoneString returns an empty optional string.
func NoneString() OptionString {
	return OptionString{
		op: option.NoneString(),
	}
}

func (m OptionString) Empty() OptionString {
	return NoneString()
}

func (m OptionString) Append(other OptionString) OptionString {
	switch {
	case m == NoneString():
		return other
	case other == NoneString():
		return m
	default:
		return SomeString(m.op.Append(other.op).UnwrapOrDefault())
	}
}

func (m OptionString) Filtered(fn func() bool) OptionString {
	if fn() {
		return m
	}

	return NoneString()
}

func (m OptionString) UnwrapOr(s string) string {
	return m.op.UnwrapOr(s)
}
