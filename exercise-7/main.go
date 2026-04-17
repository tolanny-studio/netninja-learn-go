package main

import (
	"fmt"
)

var array = []string{"Lawal", "Taoheed", "Tolani"}

func greet(name string) {
	fmt.Println("Hello", name, "!")
}

func greetClass(arr []string, f func(string)) {
	for _, _name := range arr {
		f(_name)
	}
}

func variadic(nums ...int) int {
	// variadic function, when the number of incoming argument is unknown
	total := 0

	for _, num := range nums {
		total += num
	}
	return total
}

func init() {
	// init function
	fmt.Println("I am the first init function")
}
func init() {
	// init function
	fmt.Println("I am the second init function")
}

func main() {

	// functions 

	
	// defer function
	defer greetClass(array, greet)
	result := variadic(4, 7, 9)
	fmt.Println(result)
	defer func(statement string) {
		fmt.Println(statement)
	}("I love go programming")


}
