package monoid

// String represents a string monoid.
//
// see by https://medium.com/@groveriffic/monoids-for-gophers-907175bb6165
type String string

func ForString(s string) String {
	return String(s)
}

func (m String) Empty() String {
	return ""
}

func (m String) Append(other String) String {
	return m + other
}

func (m String) Filtered(fn func() bool) String {
	if fn() {
		return m
	}

	return ""
}

func (m String) Unwrap() string {
	return string(m)
}

func (m String) UnwrapOr(s string) string {
	if m == "" {
		return s
	}

	return string(m)
}
