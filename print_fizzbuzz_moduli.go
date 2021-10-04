package gofizzbuzz

import "fmt"

func PrintFizzBuzzModuli() {
	for i := 1; i <= 15; i++ {
		fmt.Println(FizzBuzzModuli(i))
	}
}
