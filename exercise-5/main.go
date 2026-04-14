package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	// The standard library

	// strings

	greeting := "hello there friends!"

	contain_string := strings.Contains(greeting, "hello")
	fmt.Println(contain_string)

	fmt.Println(strings.ReplaceAll(greeting, "hello", "Hi"))
	fmt.Println(strings.ToUpper(greeting))

	// ints
	nums := []int{100,33,10,22, 21, 67, 89}
	sort.Ints(nums)
	fmt.Println(nums)
}
