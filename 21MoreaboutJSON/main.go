package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //aliases
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON")
	EncodeJSON()

}

func EncodeJSON() {
	lcoCourses := []course{
		{"ReactJS BootCamp", 299, "Udemy", "abc123", []string{"WebDev", "js"}},
		{"Angular BootCamp", 199, "Udemy", "a4c123", []string{"full-stack", "js"}},
		{"NodeJs BootCamp", 299, "Udemy", "ab3123", nil},
	}

	//package this data as json data

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //consumable json data
	CheckError(err)
	fmt.Printf("%s\n", finalJson)
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
