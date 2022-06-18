package main

import (
	"fmt"
)

func main() {
	calc := calculator{}
	calc.Plus(5)
	calc.Plus(7)
	calc.Multiply(4)
	calc.Minus(6)
	fmt.Println(calc.Equals())
}
