package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"` // inside json it will be called coursename not name
	Price    int
	platform string
	Password string   `json:"-"`              // password will not be shown in json
	Tags     []string `json:"tags,omitempty"` // if tags are nil it will not show in json because of omitempty
}

func main() {
	fmt.Println("more josn in golang")
	encodeJSON()
}

func encodeJSON() {
	courses := []course{
		{"Go", 50, "Udemy", "password123", []string{"Go", "Programming"}},
		{"Python", 40, "Udemy", "password123", []string{"Python", "Programming"}},
		{"Java", 30, "Udemy", "password123", nil},
	}
	fmt.Println(courses)
	// package this data as json data
	finaljson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", finaljson)
}
