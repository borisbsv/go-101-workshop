package animals

type Dog struct {
	Legs     int
	Greeting string
	Cuteness uint64
}

func (d Dog) Greet() string {
	return d.Greeting
}
