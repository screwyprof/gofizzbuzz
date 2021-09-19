package monoid

// String represents a string monoid.
//
// see by https://medium.com/@groveriffic/monoids-for-gophers-907175bb6165
type String string

func NewString(s string) String {
	return String(s)
}

func NewEmptyString() String {
	return ""
}

func (m String) Empty() String {
	return ""
}

func (m String) Append(other String) String {
	return m + other
}

func (m String) AppendIf(fn func() bool, other String) String {
	if fn() {
		return m.Append(other)
	}

	return m
}

func (m String) String() string {
	return string(m)
}
