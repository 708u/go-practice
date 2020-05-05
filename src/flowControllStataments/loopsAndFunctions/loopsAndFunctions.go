package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) (float64, int) {
	// 前回の計算結果を保持する
	// 前回の計算結果 - 今回の計算結果 = 差分
	z, diff := 1.0, 1.0
	i := 0
	for ; i < 1000; i++ {
		z -= (z*z - x) / (2 * z)

		if math.Abs(diff-z) > 0.00000001 {
			diff = z
		}
	}

	return z, i
}

func main() {
	fmt.Println(sqrt(2))
	fmt.Println(math.Sqrt(2))
}
