package main

import "fmt"

func main() {
	fmt.Println(sprit(17))
}

func sprit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}
