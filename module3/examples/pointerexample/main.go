package main

import (
	"fmt"
)

type Dog struct {
	name string
}

func rename(d Dog) {
	d.name = "Pesho"
	fmt.Printf("Renamed inside the value func: %#v\n", d)
}

func renameRef(d *Dog) {
	d.name = "Vihren"
}

func main() {
	d := Dog{name: "Stamat"}
	fmt.Printf("Original dog: %#v\n", d)

	rename(d)
	fmt.Printf("Attempted to rename passing in a value:   %#v\n", d)

	renameRef(&d)
	fmt.Printf("Attempted to rename passing in a pointer: %#v\n", d)
}
