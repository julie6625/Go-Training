package main

import (
	"errors"
	"fmt"
)

func check(err error) {
	if err != nil {
		panic("pan")
	}
}

func main() {
	defer func() {
		fmt.Println(recover()) //3.recover the panic,return the message inside panic
	}()

	fmt.Println(123)
	err := errors.New("math: square root of negative number") //1.create error
	check(err)                                                //2.trigger panic
}
