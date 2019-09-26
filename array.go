package main

import "fmt"

func main() {
	array1 := [4]string{
		"hey",
		"there",
		"you're",
		"an",
	}

	for i, e := range array1 {
		fmt.Println(i, ":", e)
	}
}
