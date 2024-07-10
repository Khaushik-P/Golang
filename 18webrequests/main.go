package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com/"

func main() {
	fmt.Println("Welcome to Web	Requests")
	response, err := http.Get(url)
	checkError(err)
	fmt.Printf("Response is of type: %T\n", response)
	defer response.Body.Close() //callers responsibility to close the connection
	databytes, err := ioutil.ReadAll(response.Body)
	checkError(err)
	fmt.Println(string(databytes))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
