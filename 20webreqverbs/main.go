package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web verbs")
	PerformGetRequest()
	PerformPostJsonRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	CheckError(err)

	defer response.Body.Close()
	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content Length is: ", response.ContentLength)

	var responseString strings.Builder
	content, err := io.ReadAll(response.Body)
	bytecnt, _ := responseString.Write(content)
	fmt.Println("Byte cnt is: ", bytecnt)
	fmt.Println(responseString.String())
	CheckError(err)
	//fmt.Println(content)
	//fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestBody := strings.NewReader(`
	{
		"coursename": "reactjs",
		"price": 0,
		"platform": "google.in"
	}
`)
	response, err := http.Post(myurl, "application/json", requestBody)
	CheckError(err)
	defer response.Body.Close()
	content, err := io.ReadAll(response.Body)
	CheckError(err)
	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myurl = "http://localhost:8000/postform"

	//form data
	data := url.Values{}
	data.Add("firstname", "Khaushik")
	data.Add("lastname", "P")
	data.Add("Email", "khau@gmail.com")
	response, err := http.PostForm(myurl, data)
	CheckError(err)
	content, err := io.ReadAll(response.Body)
	CheckError(err)
	defer response.Body.Close()
	fmt.Println(string(content))
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
