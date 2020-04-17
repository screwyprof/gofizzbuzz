package gofizbuzz

import (
	"fmt"
	"strconv"
)

func PrintFizzBuzz() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}

func FizzBuzz(n int) string {
	switch {
	case n%15 == 0:
		return "FizzBuzz"
	case n%5 == 0:
		return "Buzz"
	case n%3 == 0:
		return "Fizz"
	default:
		return strconv.Itoa(n)
	}
}
