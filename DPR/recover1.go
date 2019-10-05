package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	defer fmt.Println("defer main") // will this be called when panic?
	var user = os.Getgid()
	fmt.Println(user)
	go func() {
		defer func() {
			fmt.Println("defer caller")
			if err := recover(); err != nil {
				fmt.Println("recover success.")
			}
		}()
		func() {
			defer func() {
				fmt.Println("defer here")
			}()

			if user == -1 {
				panic("should set user gid.")
			}
			fmt.Println("after panic")
		}()
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("final?")
}
