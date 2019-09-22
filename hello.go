package main

func main() {
	var a = 50

	if a > 40 && a < 49 {
		println("right")
	} else {
		println("wrong")
	}

	if a > 40 || a < 49 {
		println("right")
	} else {
		println("wrong")
	}

	if !(a > 40 && a < 49) {
		println("right")
	} else {
		println("wrong")
	}

	if !(a > 40 || a < 49) {
		println("right")
	} else {
		println("wrong")
	}
}
