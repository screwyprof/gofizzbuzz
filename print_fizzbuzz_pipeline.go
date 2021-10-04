package gofizzbuzz

import (
	"context"
	"time"

	"github.com/screwyprof/gofizzbuzz/pipeline"
)

func PrintFizzBuzzPipeline() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	p := pipeline.NewFizzBuzz(fizzbuzz())
	p.Run(ctx, 15)
}

func fizzbuzz() func(in interface{}) interface{} {
	return func(in interface{}) interface{} {
		return FizzBuzz(int(in.(uint64)))
	}
}
