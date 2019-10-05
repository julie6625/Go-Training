package main

import "log"

type Color int

const (
	White Color = iota
	Black
	Green
	Yellow
)

type Size int

const (
	Large Size = iota
	Middle
	Small
	ExtraLarge
)

type Clothes struct {
	color Color
	size  Size
}

type Param struct {
	Color Color
	Size  Size
}

func MakeClothes(param Param) *Clothes {
	c := new(Clothes)

	c.color = param.Color
	c.size = param.Size

	return c
}

func main() {
	// Clothes with custom parameters
	c1 := MakeClothes(Param{Color: Black, Size: Middle})

	if !(c1.color == Black) {
		log.Fatal("Wrong color")
	}

	if !(c1.size == Middle) {
		log.Fatal("Wrong size")
	}

	// Clothes with default parameters
	c2 := MakeClothes(Param{})

	if !(c2.color == White) {
		log.Fatal("Wrong color")
	}

	if !(c2.size == Large) {
		log.Fatal("Wrong size")
	}
}
