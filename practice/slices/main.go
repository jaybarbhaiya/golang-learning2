package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices:")
	fmt.Println("This is a wrapper function around arrays")

	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("Array does not have a size defined, this will create a slice", colors)

	colors = append(colors, "Orange")
	fmt.Println(colors)

	colors = append(colors[2:len(colors)])
	fmt.Println(colors)

	colors = append(colors[:len(colors)-1])
	fmt.Println(colors)

	// third argument is the cap
	numbers := make([]int, 5)
	fmt.Println(numbers)

	numbers[0] = 1
	numbers[1] = 7
	numbers[2] = 4
	numbers[3] = 8
	numbers[4] = 2
	fmt.Println(numbers)

	numbers = append(numbers, 22)
	sort.Ints(numbers)
	fmt.Println(numbers)

}
