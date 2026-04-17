package main

import (
	"fmt"
)

func main() {
	// Loops

	// while loop

	// x := 0
	// for x < 5 {
	// 	fmt.Printf("The value of x is %d\n", x)
	// 	x++
	// }

	//range
	
	// for u := range 5 {
	// 	fmt.Println(u)
	// }

	names := []string{"Lawal", "Taoheed", "Tolani"}

	// using range

	for _, name := range names {
		fmt.Println(name)
	}

	// using generic for loop

	for _name := 0; _name < len(names); _name++ {
		fmt.Println(names[_name])
	}
}
