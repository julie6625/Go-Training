package main

import "fmt"

func multi(x float64, y float64) float64 {
	return x * y
}

func multiFive(x float64) float64 {
	return multi(x, 5)
}
func main() {
	fmt.Println(multi(7, 4))
	fmt.Println(multiFive(7))
}
