package option

var NilString *string

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
type String struct {
	str *string
}

func NewString(s *string) String {
	if s == nil {
		return NoneString()
	}

	return SomeString(*s)
}

func NewStringIf(fn func() bool, s string) String {
	if fn() {
		return SomeString(s)
	}

	return NoneString()
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
