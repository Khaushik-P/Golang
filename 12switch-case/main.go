package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch-case")

	// Seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// Generate a random number between 1 and 6
	dicenumber := rand.Intn(6) + 1
	fmt.Println("Value of dice number:", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 and you can open")

	case 2:
		fmt.Println("You can move 2 spot")

	case 3:
		fmt.Println("You can move 3 spot")

	case 4:
		fmt.Println("You can move 4 spot")
		fallthrough

	case 5:
		fmt.Println("You can move 5 spot")

	case 6:
		fmt.Println("You can move 6 spot and roll dice again")

	default:
		fmt.Println("What was that!")
	}
}
