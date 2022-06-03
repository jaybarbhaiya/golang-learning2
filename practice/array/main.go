package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays:")
	var colors [3]string
	colors[0] = "Red"
	colors[1] = "Blue"
	colors[2] = "Green"
	fmt.Println(colors)
	fmt.Println("First value in colors array:", colors[0])

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println(numbers)

	fmt.Println("Length of colors", len(colors))
	fmt.Println("Length of numbers", len(numbers))
}
