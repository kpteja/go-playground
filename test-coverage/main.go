package main

import "fmt"

// Size returns size.
func Size(a int) string {
	switch {
	case a < 0:
		return "negative"
	case a == 0:
		return "zero"
	case a < 10:
		return "small"
	case a < 100:
		return "big"
	case a < 1000:
		return "huge"
	}
	return "enormous"
}

func main() {
	size := Size(100)
	fmt.Println(size)

	rating := Rate(5)
	fmt.Println(rating)
}
