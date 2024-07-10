package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps lecture")

	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["RB"] = "Ruby"
	language["PY"] = "Python"

	fmt.Println("List of all languages: ", language)
	fmt.Println("JS is short form for: ", language["JS"])

	delete(language, "RB")

	fmt.Println(language)

	//for loops in golang
	for key, value := range language {
		fmt.Printf("For key %v,value is %v\n", key, value)
	}

}
