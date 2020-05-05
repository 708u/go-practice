package main

import "fmt"

func add(x, y int) int {
	return x + y
}

func times(a, b int) int {
	return a * b
}

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(times(3, 5))
}
