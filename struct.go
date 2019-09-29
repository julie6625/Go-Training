package main

import "fmt"

type address struct {
	x float64
	y float64
	z float64
}

func main() {
	p := address{x: 5.0, y: 6.0, z: 7.0}
	w := address{x: 7.0, y: 1.0, z: 4.0}

	midpoint := address{x: (p.x + w.x) / 2, y: (p.y + w.y) / 2, z: (p.z + w.z) / 2}
	fmt.Print("p點和w點的中點為")
	fmt.Println("[", midpoint.x, ",", midpoint.y, ",", midpoint.z, "]")
}
