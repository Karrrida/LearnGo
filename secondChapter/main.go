package main

import "fmt"

func main() {
	var a, b int = sumInt(1, 2, 3, 4)

	fmt.Printf("length is %d, and sum is %d", a, b)
}

func sumInt(a ...int) (lenInt, sum int) {

	for idx := range a {
		lenInt++
		sum += a[idx]
	}

	return
}

func vote(x int, y int, z int) int {
	var result int = 0

	if x+y+z > 1 {
		result = 1
	}

	return result
}
