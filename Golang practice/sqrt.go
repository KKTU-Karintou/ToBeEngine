package main

import (
	"fmt"
	"math"
)

func MySqrt(x float64) float64 {
	z, prev_z := 0.1, 0.1
	var s float64

	for i := 1; ; i++ {
		z -= (z*z - x) / (2 * z)
		fmt.Printf("LOOP:%d => z:%.10f prev_z:%.10f s:%10f\n", i, z, prev_z, s)

		s = prev_z - z
		if s < 0 {
			s *= -1
		}

		if s < 0.00001 {
			break
		}

		prev_z = z
	}

	return z
}

func main() {
	n := float64(13)

	fmt.Println("INIT:", n)
	fmt.Println("MySqrt  -> ", MySqrt(n))
	fmt.Println("library -> ", math.Sqrt(n))
}
