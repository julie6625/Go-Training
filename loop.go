package main

func main() {
	for a := 1; a <= 100; a++ {
		if a%2 == 1 {
			continue
		}

		println(a)

	}
}
