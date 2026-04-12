package main

import (
	"fmt"
)

func main() {
	// Printing and formatting strings
	name := "taoheed"
	age := 45
	height := 1.89

	fmt.Print("Hello ")
	fmt.Print("World! \n")

	fmt.Println("My name is", name, "my age is", age, "and I am", height, "tall")

	fmt.Printf("My name is %v, and I am %v of age\n", name, age)
	fmt.Printf("I am %q by name \n", name)

	whoami := fmt.Sprintf("I am %s, and I am %d by age. My height is %.3f cm", name, age, height)
	fmt.Println(whoami)

}
