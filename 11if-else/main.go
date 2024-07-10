package main

import "fmt"

func main() {
	fmt.Println("Welcome to if-else")

	logincnt := 10
	var result string

	if logincnt < 10 {
		result = "Regular cnt"
	} else if logincnt > 10 {
		result = "WatchOut"
	} else {
		result = "Exactly 10 cnt"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}

	if num := 2; num < 10 {
		fmt.Print("Number is less than 10")
	} else {
		fmt.Print("Number is greater than 10")
	}

}
