package gofizbuzz

import "strconv"

func PrintFizzBuzz() {

}

func FizzBuzz(n int) string {
	switch {
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(n)
	}
}
