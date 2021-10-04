package gofizzbuzz

import "fmt"

func PrintFizzBuzzFilter() {
	for i := 1; i <= 15; i++ {
		fmt.Println(FizzBuzzFilter(i))
	}
}
