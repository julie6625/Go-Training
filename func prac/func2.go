package main

import (
	"fmt"
)

func sum1(arr ...float64) float64 {
	sum := 0.0

	for _, e := range arr {
		sum += e
	}

	return sum
}

func main() {
	s := sum1(1, 2, 3, 4, 5)

	fmt.Println(s)
}
