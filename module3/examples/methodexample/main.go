package main

import (
	"methodexample/animals"

	"fmt"
)

func main() {
	d := animals.Dog{
		Legs:     4,
		Greeting: "Woof!",
		Cuteness: uint64(0),
	}

	// No sugar coating
	fmt.Println(animals.Dog.Greet(d))

	// Sugar coating
	fmt.Println(d.Greet())
}
