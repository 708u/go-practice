package main

import "fmt"

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
	x, y, z := superswap("I", "am", "here")
	fmt.Println(x, y, z)
}

func swap(x, y string) (string, string) {
	return y, x
}

func superswap(x, y, z string) (string, string, string) {
	return z, y, x
}
