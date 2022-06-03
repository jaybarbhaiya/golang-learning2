package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Maps:")

	states := make(map[string]string)
	fmt.Println(states)

	states["MH"] = "Maharastra"
	states["AP"] = "Andra Pradesh"
	states["GJ"] = "Gujarat"
	fmt.Println(states)

	delete(states, "AP")
	fmt.Println(states)

	states["RJ"] = "Rajasthan"
	maharastra := states["MH"]
	fmt.Println(maharastra)

	for key, value := range states {
		fmt.Printf("%v, %v\n", key, value)
	}

	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println("Sorted", keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
