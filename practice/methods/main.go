package main

import (
	"fmt"
)

func main() {
	fmt.Println("Methods")

	jay := Person{"Jay", 32, "English"}
	fmt.Println(jay)

	language := jay.Speaks()
	fmt.Println("Jay speaks:", language)

	jay.SpeakThreeTimes()
	jay.SpeakThreeTimes()
}

type Person struct {
	Name     string
	Age      int
	Language string
}

func (p Person) Speaks() string {
	return p.Language
}

func (p Person) SpeakThreeTimes() {
	p.Language = fmt.Sprintf("%v %v %v", p.Language, p.Language, p.Language)
	fmt.Println(p.Language)
}
