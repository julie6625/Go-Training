package main

import (
	"fmt"
)

func main() {
	defer func() {
		if err := recover(); err != nil { //declare err first , set err's condition
			fmt.Println("Recover from panic")
		}
	}()

	panic("Some error")

	// It didn't occur.
	fmt.Println("More message")
}
