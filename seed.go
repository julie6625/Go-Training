package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	s := rand.NewSource(time.Now().UnixNano())

	r := rand.New(s)

	n := r.Intn(10) + 1

	fmt.Println(n)
}
