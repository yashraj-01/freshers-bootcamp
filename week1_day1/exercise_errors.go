package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt2 negative number: %v", float64(e))
}

func Sqrt2(x float64) (float64, error) {
	if x < 0 {
		return 0.0, ErrNegativeSqrt(x)
	}
	return math.Sqrt(x), nil
}

func main() {
	fmt.Println(Sqrt2(2))
	fmt.Println(Sqrt2(-2))
}
