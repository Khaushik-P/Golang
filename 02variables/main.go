package main

import "fmt"

//First Letter as Capital is known as Public

const LoginToken string = "amdadada" //Public

func main() {
	var username string = "Khaushik"
	fmt.Println(username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var smallVal int = 258
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallValFloat float64 = 258.161616161616152 //float32 and 64
	fmt.Println(smallValFloat)
	fmt.Printf("Variable is of type: %T \n", smallValFloat)

	var another float32
	fmt.Println(another)
	fmt.Printf("Variable is of type: %T \n", another)

	//implicit type
	var website = "Khaushik is Pro"
	fmt.Println(website)

	//no var style

	numberofUsers := 30000
	fmt.Println(numberofUsers)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type: %T \n", LoginToken)
}
