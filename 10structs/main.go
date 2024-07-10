package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	khaushik := User{"Khaushik", "khau@gmail.com", true, 20}
	fmt.Println(khaushik)
	fmt.Printf("Khaushik's details are: %+v\n", khaushik)
	fmt.Printf("Name is %v and email is %v.", khaushik.Name, khaushik.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
