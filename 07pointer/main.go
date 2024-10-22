package main

import "fmt"

func main() {
	fmt.Println("using pointer")

	var ptr *int
	fmt.Println("value of ptr is:", ptr)    //value of ptr is: <nil>
	fmt.Println("address of ptr is:", &ptr) //address of ptr is: 0xc

	mynumber := 20
	ptr = &mynumber

	fmt.Println("value is:", *ptr) //value is: 20
	fmt.Println("value is:", ptr)  // value is: 0xc00008c098 -- mem location of 20
	fmt.Println("value is:", &ptr) // value is: 0xc000090050 -- mem location of ptr

	*ptr += 20
	fmt.Println("new value:", *ptr) //new value: 40

}
