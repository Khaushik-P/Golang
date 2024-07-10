package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study of golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) //standard formatting for time check documentation

	createDate := time.Date(2021, time.August, 12, 23, 23, 0, 0, time.UTC)
	fmt.Println(createDate)
	fmt.Println(createDate.Format("01-02-2006 Monday"))

	//create exe with go build command
	//For specific OS use GOOS="OS-NAME" go build {GOOS is known from the command go env}
}
