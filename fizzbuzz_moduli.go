package gofizzbuzz

import (
	"strconv"

	"github.com/screwyprof/gofizzbuzz/fastdiv"
)

var filters = []struct {
	word string
	div  uint64
}{
	{
		word: "Fizz",
		div:  fastdiv.M3,
	},
	{
		word: "Buzz",
		div:  fastdiv.M5,
	},
}

func FizzBuzzModuli(n int) string {
	var res string

	for _, f := range filters {
		if fastdiv.IsDivisible(uint64(n), f.div) {
			res += f.word
		}
	}

	if res != "" {
		return res
	}

	return strconv.FormatInt(int64(n), 10)
}
