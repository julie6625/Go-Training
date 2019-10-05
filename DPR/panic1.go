package main

import (
	"fmt"
)

func main() {
	defer func() {
		fmt.Println(recover())
	}()

	fmt.Println(123)
	panic("Panic!")
	fmt.Println(456)

}
