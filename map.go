package main

import (
	"fmt"
	"log"
)

func main() {
	m := make(map[string]string)

	m["thresh"] = "hook"
	m["nami"] = "splash"
	m["bard"] = "meeps"
	m["blitz"] = "grab"

	v, ok := m["thresh"]
	if !ok {
		log.Fatal("it should be true")
	}

	if !(v == "hook") {
		log.Fatal("Wrong value")
	}

	_, ok = m["sona"]
	if !(ok == false) {
		log.Fatal("it should be false")
	}

	for i := 0; i < 10; i++ {
		for k, v := range m {
			fmt.Println(fmt.Sprintf("%s: %s", k, v))
		}

		fmt.Println("")
	}
	if !(m["sona"] == "") {
		log.Fatal("no such thing2")
	}
}
