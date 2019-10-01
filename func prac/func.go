package main

import "fmt"

func sum(a int, b int) int {
	return a + b
}

func print2() {
	q := 1
	w := 5
	s := sum(q, w)
	fmt.Printf("%d + %d = %d", q, w, s)
}
func main() {
	print2()
}
