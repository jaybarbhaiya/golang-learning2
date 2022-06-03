package main

import (
	"fmt"
)

func main() {
	fmt.Println("Struct")

	ashwini := Person{"Ashwini", 26, "Baapla"}
	fmt.Println(ashwini)

	fmt.Printf("%+v\n", ashwini)
	fmt.Printf("Name:%v\nAge:%v\nPetname:%v\n", ashwini.Name, ashwini.Age, ashwini.Petname)
}

type Person struct {
	Name    string
	Age     int
	Petname string
}
