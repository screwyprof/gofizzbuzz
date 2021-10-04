package gofizzbuzz

import (
	"fmt"

	"github.com/genkami/dogs/types/iterator"
)

func PrintFizzBuzzFunctional() {
	it := iterator.Map(iterator.Range[int](1, 100), FizzBuzzFunctional)
	iterator.ForEach(it, func(s string) { fmt.Println(s) })
}
