package option

// SomeString wraps the s into a String.
func SomeString(s string) String {
	return String{&s}
}

// NoneString returns an empty optional string.
func NoneString() String {
	return String{}
}

// String represents Option<String>.
//
// Option X = *X
// None = nil
// Some x = &x
//
// see https://github.com/leighmcculloch/go-optional/blob/master/string_generated.go
type String struct {
	str *string
}

func FromStringPtr(s *string) String {
	if s == nil {
		return NoneString()
	}

	return SomeString(*s)
}

func (s String) Append(other String) String {
	switch {
	case s == NoneString():
		return other
	case other == NoneString():
		return s
	default:
		return SomeString(s.UnwrapOrDefault() + other.UnwrapOrDefault())
	}
}

func (s String) UnwrapOrDefault() string {
	if s.str == nil {
		return ""
	}

	return *s.str
}

func (s String) UnwrapOr(val string) string {
	if s.str == nil {
		return val
	}

	return *s.str
}
