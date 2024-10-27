package main

import (
	"encoding/json" // Imports the JSON package for encoding and decoding JSON
	"fmt"
)

// Define the structure (schema) for a course
type course struct {
	Name     string `json:"coursename"` // `Name` will be renamed to "coursename" in JSON
	Price    int
	Platform string
	Password string   `json:"-"`              // Password will be hidden in JSON (the "-" option omits this field)
	Tags     []string `json:"tags,omitempty"` // `Tags` field is renamed to "tags" and will be omitted if empty
}

func main() {
	fmt.Println("more json in golang")
	encodeJSON() // Call the function to encode data to JSON
	decodeJSON() // Call the function to decode data from JSON
}

// Function to convert Go data into JSON format
func encodeJSON() {
	// Create a slice of course objects with different data
	courses := []course{
		{"Go", 50, "Udemy", "password123", []string{"Go", "Programming"}},
		{"Python", 40, "Udemy", "password123", []string{"Python", "Programming"}},
		{"Java", 30, "Udemy", "password123", nil},
	}

	fmt.Println(courses)

	// Convert `courses` to JSON format with indentation for readability
	finaljson, err := json.MarshalIndent(courses, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finaljson)
}

// Function to decode (convert) JSON data into Go struct data
func decodeJSON() {
	// Define a JSON string with course data in it (using `[]byte` to handle raw JSON data)
	jsondata := []byte(`[
			{
				"coursename": "Go",
				"Price": 50,
				"tags": [
					"Go",
					"Programming"
				]
			},
			{
				"coursename": "Python",
				"Price": 40,
				"tags": [
					"Python",
					"Programming"
				]
			}
		]`)
	// Define a variable to hold the decoded course data
	var allsourses course

	// Check if the JSON data is valid and can be parsed
	checkvalid := json.Valid(jsondata)

	if checkvalid {
		// If JSON is valid, print confirmation
		fmt.Println("JSON was valid")

		// Decode (unmarshal) the JSON data into the `allsourses` variable
		json.Unmarshal(jsondata, &allsourses)
		fmt.Printf("%#v\n", allsourses) // Print the struct with its field names
	} else {
		// If JSON is not valid, print an error message
		fmt.Println("JSON was invalid")
	}

	var genericData []map[string]interface{}
	err := json.Unmarshal(jsondata, &genericData)
	if err != nil {
		fmt.Println("Error during unmarshalling to generic type:", err)
	} else {
		fmt.Printf("%#v\n", genericData)
	}
}
