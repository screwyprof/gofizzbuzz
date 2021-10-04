package fastdiv

var (
	M3  = 1 + uint64(0xffffffffffffffff)/3
	M5  = 1 + uint64(0xffffffffffffffff)/5
	M15 = 1 + uint64(0xffffffffffffffff)/15
)

// IsDivisible checks whether n % m == 0, given precomputed m.
//
// https://lemire.me/blog/2016/06/27/a-fast-alternative-to-the-modulo-reduction/
func IsDivisible(n, m uint64) bool {
	return n*m <= m-1
}
