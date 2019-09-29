package main

import (
	"fmt"
	"math"
)

type coordinate struct {
	x float64
	y float64
}

type area struct {
	con float64
	p   float64
}

func main() {
	a := coordinate{1.0, 2.0}
	b := coordinate{5.0, 14.0}
	c := coordinate{13.0, 4.0}

	triangle := area{con: 0.5, p: ((a.x * b.y) + (b.x * c.y) + (c.x * a.y)) - ((b.x * a.y) + (c.x * b.y) + (a.x * c.y))}

	fmt.Println(math.Abs(triangle.p))
	fmt.Println(math.Abs(triangle.con))

	trianglearea := triangle.con * triangle.p
	fmt.Println("三角形面積為", math.Abs(trianglearea))
}
