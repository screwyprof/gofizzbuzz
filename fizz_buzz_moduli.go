package gofizbuzz

import "strconv"

var moduli = map[int]string{
	3: "Fizz",
	5: "Buzz",
}

func FizzBuzzModuli(n int) string {
	var res string

	for m, v := range moduli {
		if n%m == 0 {
			res += v
		}
	}

	if res != "" {
		return res
	}

	return strconv.Itoa(n)
}
