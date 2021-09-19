package monoid

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

func (m String) String() string {
	return string(m)
}
