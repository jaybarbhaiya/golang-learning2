package main

import (
	"fmt"
)

func main() {
	fmt.Println("for Statement")

	colors := []string{"Red", "Green", "Blue"}
	fmt.Println(colors)

	for i := 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}

	fmt.Println("----------------------")
	for i := range colors {
		fmt.Println(colors[i])
	}

	fmt.Println("----------------------")
	for _, color := range colors {
		fmt.Println(color)
	}

	fmt.Println("----------------------")
	value := 1
	for value < 10 {
		fmt.Println("Value:", value)
		value++
	}

	fmt.Println("----------------------")
	sum := 1
	for sum < 1000 {
		sum += sum
		fmt.Println("Sum=", sum)
		if sum > 200 {
			goto theEnd
		}
	}

theEnd:
	fmt.Println("End of Program")
}
