package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers class")

	// var ptr *int

	// fmt.Println("Value of Pointer is:", ptr)

	myNumber := 23

	var ptr = &myNumber //reference

	fmt.Println("Value of Pointer is:", ptr)  //ptr normally is used to find the address of the value
	fmt.Println("Value of Pointer is:", *ptr) //* is used to see what is inside the pointer

	*ptr = *ptr + 2

	fmt.Println("New value of pointer is:", myNumber)

}
