package gofizbuzz

import "strconv"

var filters = []struct {
	div  int
	word string
}{
	{
		div:  3,
		word: "Fizz",
	},
	{
		div:  5,
		word: "Buzz",
	},
}

func FizzBuzzModuli(n int) string {
	var res string

	for _, f := range filters {
		if n%f.div == 0 {
			res += f.word
		}
	}

	if res != "" {
		return res
	}

	return strconv.Itoa(n)
}
