package gofizzbuzz_test

import (
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"testing/quick"

	"github.com/screwyprof/gofizzbuzz"
)

func TestFizzBuzzProperties(t *testing.T) {
	t.Parallel()

	fizzBuzzers := []func(n int) string{
		gofizzbuzz.FizzBuzz,
		gofizzbuzz.FizzBuzzFilter,
		gofizzbuzz.FizzBuzzModuli,
		gofizzbuzz.FizzBuzzFunctional,
	}

	for _, fizzBuzzer := range fizzBuzzers {
		checkFizzBuzzProperties(t, fizzBuzzer)
	}
}

func checkFizzBuzzProperties(t *testing.T, fizzBuzzer func(n int) string) {
	mul3Fizz := func(i int) bool {
		result := fizzBuzzer(i * 3)
		return strings.HasPrefix(result, "Fizz")
	}

	mul5Buzz := func(i int) bool {
		result := fizzBuzzer(i * 5)
		return strings.HasSuffix(result, "Buzz")
	}

	restAreNums := func(i int) bool {
		result := fizzBuzzer(i)
		return result == strconv.FormatInt(int64(i), 10)
	}

	t.Run("it should start with Fizz for all multiples of 3", func(t *testing.T) {
		t.Parallel()
		if err := quick.Check(mul3Fizz, genCfg(div3Generator)); err != nil {
			t.Errorf("expected Fizz for number: %v", err)
		}
	})

	t.Run("it should end with Buzz for all multiples of 5", func(t *testing.T) {
		t.Parallel()
		if err := quick.Check(mul5Buzz, genCfg(div5Generator)); err != nil {
			t.Errorf("expected Buzz for number: %v", err)
		}
	})

	t.Run("it should return number as string for all values non-divisible by 3 or 5", func(t *testing.T) {
		t.Parallel()
		if err := quick.Check(restAreNums, genCfg(non3or5DivGenerator)); err != nil {
			t.Errorf("expected string number for number: %v", err)
		}
	})
}

func div3Generator(r *rand.Rand) int {
	return r.Intn(math.MaxInt32 / 3)
}

func div5Generator(r *rand.Rand) int {
	return r.Intn(math.MaxInt32 / 5)
}

func non3or5DivGenerator(r *rand.Rand) int {
	var num int
	for {
		num = r.Intn(math.MaxInt32)
		if !(num%3 == 0 || num%5 == 0) {
			break
		}
	}
	return num
}

func genCfg(numGen func(*rand.Rand) int) *quick.Config {
	f := func(args []reflect.Value, r *rand.Rand) {
		for i := range args {
			num := numGen(r)
			args[i] = reflect.ValueOf(num)
		}
	}
	return &quick.Config{Values: f}
}
