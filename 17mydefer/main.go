package main

import (
	"fmt"
)

func printdefer() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}
}

func main() {
	// when writting this line below hello then we are getting hello world and if we are using it above hello then also it is giving hello world because when we use defer then the line perform at the last of the program

	// defer works on last in first out
	defer fmt.Println("world")
	defer fmt.Println("world1")
	defer fmt.Println("world2")
	fmt.Println("hello")
	fmt.Println("hii")
	printdefer()
}
