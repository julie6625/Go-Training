package main

import (
	"fmt"
)

func main() {
	slice1 := [][]float64{
		[]float64{1, 2, 3},
		[]float64{4, 5, 6},
	}

	for i := 0; i < 2; i++ {
		for e := 0; e < 3; e++ {
			fmt.Println(slice1[i][e])
		}
	}
	//if !(slice1[] == "bard"){
	//	log.Fatal("wrong value")
	//}
}
