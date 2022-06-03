package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	i1, i2, i3 := 34, 53, 66
	intSum := i1 + i2 + i3
	fmt.Println("Integer sum:", intSum)

	f1, f2, f3 := 23.7, 4.5, 6.3
	floatSum := f1 + f2 + f3
	fmt.Println("Float sum:", floatSum)

	floatSum = math.Round(floatSum*100) / 100
	fmt.Println("Float rounded sum:", floatSum)

	circleRadius := 14.5
	circumference := circleRadius * 2 * math.Pi
	fmt.Printf("Circumference: %.2f\n", circumference)

	n := time.Now()
	fmt.Println("The time now is:", n)

	t := time.Date(2009, time.November, 10, 21, 0, 0, 0, time.UTC)
	fmt.Println("date:", t)
	fmt.Println(t.Format(time.ANSIC))

	parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 21:00:00 2009")
	fmt.Printf("The type of parsedTime is %T\n", parsedTime)
}
