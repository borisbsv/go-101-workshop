package main

import "fmt"

var characters = []string{
	"Luke Skywalker",
	"Leia Organa",
	"Han Solo",
	"Luke Skywalker",
	"Luke Skywalker",
	"Leia Organa",
	"Yoda",
	"Yoda",
	"Leia Organa",
	"Luke Skywalker",
	"Darth Vader",
	"Darth Sidious",
}

func main() {
	var result map[string]int = count(characters)

	assert(result, "Luke Skywalker", 4)
	assert(result, "Leia Organa", 3)
	assert(result, "Han Solo", 1)
	assert(result, "Yoda", 2)
	assert(result, "Darth Vader", 0)
	assert(result, "Darth Sidious", 0)

	fmt.Println("SUCCESS")
}

func assert(result map[string]int, name string, count int) {
	res := result[name]
	if res != count {
		panic(fmt.Sprintf("ERROR: Got %d for [%s], expected %d", res, name, count))
	}
}
