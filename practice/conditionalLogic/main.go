package main

import (
	"fmt"
)

func main() {
	fmt.Println("Conditional logic")

	num := 22
	var res string

	if num < 0 {
		res = "Less than Zero"
	} else if num == 0 {
		res = "Equal to Zero"
	} else {
		res = "Greater than Zero"
	}
	fmt.Println(res)
}
