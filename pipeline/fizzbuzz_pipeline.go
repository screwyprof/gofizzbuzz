package pipeline

import (
	"context"
	"fmt"
	"sync"
)

type FizzBuzzProcessorFn = MapFn

type Pipeline struct {
	fizzBuzzer FizzBuzzProcessorFn
	wg         *sync.WaitGroup
}

func NewFizzBuzz(fizzBuzzer FizzBuzzProcessorFn) *Pipeline {
	var wg sync.WaitGroup

	return &Pipeline{wg: &wg, fizzBuzzer: fizzBuzzer}
}

func (p *Pipeline) Run(ctx context.Context, num uint64) {
	counter := func() func() interface{} {
		var counter uint64

		return func() interface{} {
			counter++

			return counter
		}
	}

	pipe := p.repeatFnTimes(ctx, num, counter())
	pipe = p.mapFn(ctx, p.fizzBuzzer, pipe)

	for v := range pipe {
		fmt.Printf("%s\n", v)
	}

	p.wg.Wait()
}

func (p *Pipeline) repeatFnTimes(ctx context.Context, times uint64, fn func() interface{}) <-chan interface{} {
	valueStream := make(chan interface{})

	p.wg.Add(1)

	go func() {
		defer p.wg.Done()
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

type MapFn func(in interface{}) interface{}

func (p *Pipeline) mapFn(ctx context.Context, fn MapFn, in <-chan interface{}) <-chan interface{} {
	out := make(chan interface{})

	p.wg.Add(1)

	go func() {
		defer p.wg.Done()
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
