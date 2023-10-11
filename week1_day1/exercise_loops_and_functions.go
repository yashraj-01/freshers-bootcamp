package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64, zInitial float64) (zBy10Iterations float64, zByNoChange float64, iterationsInNoChange int32, zByStandardLib float64) {
	zBy10Iterations, zByNoChange = zInitial, zInitial
	iterationsInNoChange = 0

	for i := 0; i < 10; i++ {
		zBy10Iterations -= (zBy10Iterations*zBy10Iterations - x) / (2 * zBy10Iterations)
	}

	for math.Abs(zByNoChange*zByNoChange-x) > 1e-16 {
		iterationsInNoChange++
		zByNoChange -= (zByNoChange*zByNoChange - x) / (2 * zByNoChange)
	}

	zByStandardLib = math.Sqrt(x)

	return
}

func main() {
	x, zInitial := 99.0, 1.0
	zBy10Iterations, zByNoChange, iterationsInNoChange, zByStandardLib := Sqrt(x, zInitial)
	fmt.Println("Z By 10 iterations:", zBy10Iterations)
	fmt.Println("Z By No Change:", zByNoChange, "and no. of iterations:", iterationsInNoChange)
	fmt.Println("Z By Standard Lib:", zByStandardLib)
}
