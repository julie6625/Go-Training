package main

import "fmt"

func main() {
	var a, b = 1, 1
	var cond = (a+b == 3)
	if cond {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
