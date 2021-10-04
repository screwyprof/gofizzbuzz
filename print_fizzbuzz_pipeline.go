package gofizzbuzz

import (
	"context"
	"fmt"
	"time"
)

func PrintFizzBuzzPipeline() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	p := repeatFnTimes(ctx, 100, counter())
	p = mapFn(ctx, fizzbuzz(), p)

	for res := range p {
		fmt.Printf("%s\n", res.(string))
	}

	<-ctx.Done()
}

func counter() func() interface{} {
	var counter uint64

	return func() interface{} {
		counter++

		return counter
	}
}

func fizzbuzz() func(in interface{}) interface{} {
	return func(in interface{}) interface{} {
		return FizzBuzz(int(in.(uint64)))
	}
}

func repeatFnTimes(ctx context.Context, times uint64, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})

	go func() {
		defer close(valueStream)

		for i := uint64(0); i < times; i++ {
			select {
			case <-ctx.Done():
				return
			case valueStream <- fn():
			}
		}
	}()

	return valueStream
}

func mapFn(ctx context.Context, fn MapFn, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	go func() {
		defer close(out)

		for v := range in {
			select {
			case <-ctx.Done():
				return
			case out <- fn(v):
			}
		}
	}()

	return out
}
