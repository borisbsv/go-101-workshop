package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type Number interface {
	constraints.Integer | constraints.Float
}

type Calc[T Number] struct {
	total T
}

func main() {
	c := &Calc[float64]{}
	c.Add(3.2)
	c.Multiply(2.5)
	c.Minus(5)
	c.Divide(2)

	fmt.Println(c.Equals())
}
