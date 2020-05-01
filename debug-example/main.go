package main

import (
	"fmt"
)

func max(arr []int) int {
	max := arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}

	return max
}

func main() {
	arr := []int{101, 95, 10, 188, 100}
	max := max(arr)
	fmt.Printf("Max element is %d\n", max)
}
