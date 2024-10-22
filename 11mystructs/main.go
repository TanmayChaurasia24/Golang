package main

import "fmt"

func main() {
	fmt.Println("using structs:")
	tanmay := User{"tanmay", "t@gmail.com", true, 21}

	fmt.Println(tanmay)       // {tanmay t@gmail.com true 21}
	fmt.Println(tanmay.Name)  // tanmay
	fmt.Println(tanmay.Email) // t@gmail.com

}

// first letter of name and ... is capital, so that it can be access by anybody
type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
