package gofizbuzz

import "strconv"

func PrintFizzBuzz() {

}

func FizzBuzz(n int) string {
	if n % 5 == 0 {
		return "Buzz"
	}
	if n % 3 == 0 {
		return "Fizz"
	}
	return strconv.Itoa(n)
}
