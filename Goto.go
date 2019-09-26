package main

func main() {
	for a := 1; a <= 100; a++ {
		if a > 30 {
			goto END
		}
		println(a)
	}

END:

	b := 1
loop:
	for b < 100 {
		if b%2 == 1 {
			b++
			goto loop
		}
		println(b)
		b++
	}
}
