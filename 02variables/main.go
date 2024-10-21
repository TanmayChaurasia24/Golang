package main

import "fmt"

// jwt := 4323232 // walrus operator is not allowed globaly
// var jwt = 342343 // no error

const Signedup string = "sadsadksajdkasjd" // this is public variable because first letter S is capital

func main() {
	var tempint int
	var tempfloat float64
	var tempstring string
	var tempbool bool
	tempint = 42
	tempfloat = 1.2
	tempstring = "tanmay kumar chaurasia"
	tempbool = false

	fmt.Printf("type of tempint is %T \n", tempint)
	fmt.Printf("type of tempfloat is %T \n", tempfloat)
	fmt.Printf("type of tempstring is %T \n", tempstring)
	fmt.Printf("type of tempbool is %T \n", tempbool)

	fmt.Println(tempint)
	fmt.Println(tempfloat)
	fmt.Println(tempstring)
	fmt.Println(tempbool)

	// implicit type

	var isloggedin = true // datatype is decided by lexer
	fmt.Println(isloggedin)
	// isloggedin = "tan" // error

	// no var style
	username := "tanmay kumar"
	fmt.Println(username)
	username = "tanmay"
	fmt.Println(username)
}
