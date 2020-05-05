package main

import "fmt"

func main() {
	fmt.Println(add(1, 3))
	fmt.Println(divide(6, 2))
}

func add(x int, y int) int {
	return x + y
}

func divide(x int, y int) int {
	return x / y
}
