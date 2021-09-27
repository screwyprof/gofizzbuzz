package monoid

// All represents a conjunction boolean monoid.
type All bool

func ForBool(b bool) All {
	return All(b)
}

func (m All) Empty() All {
	return true
}

func (m All) Append(other All) All {
	return m && other
}
