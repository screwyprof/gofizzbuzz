package monoid

//type Predicate func(in interface{}) bool
//
//func ForPredicate(fn func(in interface{}) bool) Predicate {
//	//if fn == nil {
//	//	return func(in interface{}) All {
//	//		return true // ?
//	//	}
//	//}
//
//	m := func(in interface{}) All {
//		return ForBool(fn(in))
//	}
//
//	return m
//}
//
//func (m Predicate) Empty() Predicate {
//	return nil
//}
//
//func (m Predicate) Append(other Predicate) Predicate {
//	if other == nil {
//		return m
//	}
//
//	return func(in interface{}) All {
//		return m(other(in))
//	}
//}
//
