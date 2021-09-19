package gofizbuzz

import "fmt"

func PrintFizzBuzzModuli() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzzModuli(i))
	}
}
