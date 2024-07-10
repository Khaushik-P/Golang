package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in golang")

	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	fruitlist[3] = "mango"

	fmt.Println("FruitList is:", fruitlist)
	fmt.Println("FruitList is:", len(fruitlist))

	var veglist = [5]string{"potato", "tomato", "mushroom"}

	fmt.Println("Veggie List is :", veglist)
	fmt.Println("Veggie List is :", len(veglist))

}
