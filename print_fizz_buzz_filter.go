package gofizzbuzz

import "context"

func PrintFizzBuzzFilter() {
	num := uint64(100)

	fizzBuzzer := func(in interface{}) interface{} {
		return FizzBuzz(int(in.(uint64)))
	}

	p := NewPipeline(fizzBuzzer)
	p.Run(context.Background(), num)
}
