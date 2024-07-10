package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to golang")

	content := "This needs to go in a file"

	file, err := os.Create("./myfile.txt")
	checkError(err)
	// if err != nil {
	// 	panic(err)
	// }

	length, err := io.WriteString(file, content)
	checkError(err)

	fmt.Println("Length is: ", length)
	defer file.Close()
	readfile("./myfile.txt")
}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	checkError(err)

	fmt.Println("Text data inside the file is: ", string(databyte))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
