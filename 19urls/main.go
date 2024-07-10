package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://lco.dev:3000/learn?coursename=reactjs&paymentid=qwerty"

func main() {
	fmt.Println("Welcome to handling urls in golang")
	fmt.Println(myurl)

	//parsing

	result, err := url.Parse(myurl)
	checkError(err)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.RawQuery)
	fmt.Println(result.Port())

	qparams := result.Query()
	fmt.Printf("The type of qparams is: %T\n", qparams)
	fmt.Println(qparams["coursename"])

	for _, val := range qparams {
		fmt.Println("Param is :", val)
	}

	partsOfUrl := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "learn",
		RawQuery: "coursework=reactjs",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
