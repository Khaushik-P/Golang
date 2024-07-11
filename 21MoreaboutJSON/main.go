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
	DecodeJson()

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

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJS BootCamp",
                "Price": 299,
                "website": "Udemy",
                "tags": ["WebDev","js"]
    }
	`)

	var lcocourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON is valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for key, val := range myOnlineData {
		fmt.Printf("Key is %v Value is %v and type is: %T\n", key, val, val)
	}
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}
