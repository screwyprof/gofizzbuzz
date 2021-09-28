package gofizzbuzz

import "fmt"

func PrintFizzBuzz() {
	for i := 1; i <= 100; i++ {
		fmt.Println(FizzBuzz(i))
	}
}
