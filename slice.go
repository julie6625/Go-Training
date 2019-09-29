package main

import "fmt"

func main() {
	slice1 := []int{0, 1, 2, 3, 4, 5}

	slice1 = append(slice1[0:2], slice1[3:6]...)

	for _, e := range slice1 {
		fmt.Print(e, " ")
	}
	//if !(slice1[] == "bard"){
	//	log.Fatal("wrong value")
	//}
}
