package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	for _, days := range days {
		fmt.Printf("index is  and value is %v\n", days)
	}

	rougeValue := 1

	for rougeValue < 10 {
		if rougeValue == 2 {
			goto lco
		}
		if rougeValue == 5 {
			rougeValue++
			continue
		}
		fmt.Println("Value is: ", rougeValue)
		rougeValue++
	}
lco:
	fmt.Print("Golang is fun")
}
