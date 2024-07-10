package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices")

	var fruitlist = []string{"Apple", "Tomato", "Orange"}

	fmt.Printf("Type of fruitList is %T\n", fruitlist)

	fruitlist = append(fruitlist, "Mango", "Banana")

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[:3]) //[1:],[1:3] 3 in not inclusive

	fmt.Println(fruitlist)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 945
	highScores[2] = 465
	highScores[3] = 867
	//highScores[4] = 899 [error index out of bound]

	highScores = append(highScores, 554, 666, 555)

	fmt.Println(highScores)

	sort.Ints(highScores)

	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//how to remove values from slices based on index

	var courses = []string{"reactjs", "angular", "python", "ruby", "swift"}

	var index int = 2

	courses = append(courses[:index], courses[index+1:]...)

	fmt.Println(courses)

}
