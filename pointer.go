package main

import "fmt"

func zero(a *int) {
	*a = 6
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x)

	nPtr := new(int)
	*nPtr = 2
	fmt.Print(nPtr)
}
