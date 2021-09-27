package monoid_test

//func TestPredicate(t *testing.T) {
//	t.Parallel()
//
//	// A set S equipped with a binary operation S × S → S,
//	// which we will denote •, is a monoid if it satisfies the following two axioms:
//	t.Run("it has valid identity", func(t *testing.T) {
//		t.Parallel()
//
//		//pred := func(in interface{}) bool {
//		//	t.Logf("predicated called: %v\n", in)
//		//
//		//	return true
//		//}
//
//		/*
//		public void CombineHasIdentity(Func<Guid, int> f, Guid guid)
//		{
//		    Assert.Equal(f(guid), Combine(FuncIdentity, f)(guid));
//		    Assert.Equal(f(guid), Combine(f, FuncIdentity)(guid));
//		}
//		 */
//
//
//		//There exists an element e in S such that for every element a in S,
//		//the equations e • a = a and a • e = a hold.
//		//monoidtest.AssertEqual(t, m, m.Append(m.Empty()))
//		//monoidtest.AssertEqual(t, m.Empty().Append(m), m.Append(m.Empty()))
//	})
//
//	t.Run("it has valid associativity", func(t *testing.T) {
//		t.Parallel()
//
//		a := monoid.ForBool(true)
//		b := monoid.ForBool(true)
//		c := monoid.ForBool(true)
//
//		// For all a, b and c in S, the equation (a • b) • c = a • (b • c) holds.
//		monoidtest.AssertEqual(t, a.Append(b).Append(c), a.Append(b.Append(c)))
//	})
//}
