package main

import (
	"fmt"
)

func main() {
	var age int
	age = 42
	fmt.Println("My age is", age)
	var ptr *int
	ptr = &age
	fmt.Println("The value of age is", *ptr)
	*ptr = 21
	fmt.Println("The value of age is", age)
	fmt.Println("The address of age is", &age)
}
