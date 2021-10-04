package gofizzbuzz_test

import "github.com/screwyprof/gofizzbuzz"

func ExamplePrintFizzBuzz() {
	gofizzbuzz.PrintFizzBuzz()

	// output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func ExamplePrintFizzBuzzFilter() {
	gofizzbuzz.PrintFizzBuzzFilter()

	// output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func ExamplePrintFizzBuzzFunctional() {
	gofizzbuzz.PrintFizzBuzzFunctional()

	// output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func ExamplePrintFizzBuzzModuli() {
	gofizzbuzz.PrintFizzBuzzModuli()

	// output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}

func ExamplePrintFizzBuzzPipeline() {
	gofizzbuzz.PrintFizzBuzzPipeline()

	// output:
	// 1
	// 2
	// Fizz
	// 4
	// Buzz
	// Fizz
	// 7
	// 8
	// Fizz
	// Buzz
	// 11
	// Fizz
	// 13
	// 14
	// FizzBuzz
}
