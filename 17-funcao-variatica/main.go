package main

import "fmt"

// recebe N ints para soma
func soma(nums ...int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total

}

func main() {
	fmt.Println(soma(1, 2, 3, 5, 6, 10, 100, 2000, 310, 10))

}
