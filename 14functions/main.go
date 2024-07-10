package main

import "fmt"

func main() {
	fmt.Println("Functions in golang")
	greeter()
	result := adder(2, 5)
	fmt.Println("Result:", result)
	proResult, myMessage := proAdder(2, 3, 4, 5)
	fmt.Println(proResult)
	fmt.Println(myMessage)
}

func greeter() {
	fmt.Println("Namastey from golang")
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}
	return total, "Hi"
}
