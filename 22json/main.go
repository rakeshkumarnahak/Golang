package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` //The json:"coursename" is the alias which is used as the key while creating the json
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`              //- inside the json alias makes sure that the password is not shown in the json
	Tags     []string `json:"tags,omitempty"` //omitempty property makes sure that the key value pair is not shown if its value is nil
}

func main() {
	fmt.Println("Welcome to JSON in golang")

	fmt.Println("***************** Encoding the JSON *****************")
	EncodingJSON()
	fmt.Println("***************** Dencoding the JSON *****************")
	DecodingJSON()
}

func EncodingJSON() {
	//making slice using the course struct

	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncodeonline.com", "abc123", []string{"web-dev", "js"}},
		{"MERN Bootcamp", 199, "Learncodeonline.com", "bcd123", []string{"full-stack", "js"}},
		{"Angular Bootcamp", 299, "Learncodeonline.com", "efg123", nil},
	}

	//Packaging this data as JSON data
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t") //MarchalIndent() takes 3 arguments i.e 1. data that is to be packages into JSON 	2. prefx (what should be before every line of the json) 	3. Indent (how the data inside the json is separated)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)
}

func DecodingJSON() {
	jsonDataFromWeb := []byte(`{
			"coursename": "MERN Bootcamp",
            "Price": 199,
            "website": "Learncodeonline.com",
            "tags": ["full-stack", "js"]
		}`)

	var lcoCourse course //Using coourse struct

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON data is valid")
		json.Unmarshal(jsonDataFromWeb, &lcoCourse) //json.Unmarshal(jsonData to be decoded, variable to store the decoded data(need to pass a reference because we don't want to store it in a copy of the variable ))
		fmt.Printf("%#v\n", lcoCourse)
	} else {
		fmt.Println("JSON IS NOT VALID!")
	}

	//This is the normal way of storing decoded json data in a struct but if in some case we want to store the data as a key value pair then we can follow the below method
	var myOnlineCourse map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineCourse)
	fmt.Printf("%#v\n", myOnlineCourse)

	//Looping through the key value pair
	for key, value := range myOnlineCourse {
		fmt.Printf("Key is %v and value is %v and type is %T\n", key, value, value)
	}
}
