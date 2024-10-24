package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) Getstatus() {
	fmt.Println("is user active", u.Status)
}

func (u User) Newemail() {
	u.Email = "tmm@gmail.com"
	fmt.Println("the new email value is:", u.Email)
}

func main() {
	fmt.Println("using methods in golang")

	tanmay := User{"tanmay kumar chaurasia", "t@gmail.com", true, 21}
	fmt.Println(tanmay)
	tanmay.Getstatus()
	tanmay.Newemail()   // real email does not change
	fmt.Println(tanmay) // {tanmay kumar chaurasia t@gmail.com true 21}

}
