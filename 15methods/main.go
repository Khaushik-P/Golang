package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")

	khaushik := User{"Khaushik", "khau@gmail.com", true, 20, 25}
	fmt.Println(khaushik)
	fmt.Printf("Khaushik's details are: %+v\n", khaushik)
	fmt.Printf("Name is %v and email is %v.\n", khaushik.Name, khaushik.Email)
	khaushik.GetStatus()
	khaushik.GetEmail()
	fmt.Printf("Name is %v and email is %v.\n", khaushik.Name, khaushik.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	OneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) GetEmail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of the user is:", u.Email)
}
