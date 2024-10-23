package main

import "fmt"

func greeter() {
	fmt.Println("hello sir")
}

// return type is int
func adder(val1 int, val2 int) int {
	return val1 + val2
}

// you can pass as many values you want and there return type is int and string, can return multiple values
func proadder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "done the calculation"
}

func main() {
	fmt.Println("functions in golang")
	greeter()

	result := adder(3, 5)
	fmt.Println("result is:", result)

	prores, conclusion := proadder(1, 3, 4, 5, 6, 2, 2, 32, 42, 3, 23, 23, 343)
	// prores, _ := proadder(1, 3, 4, 5, 6, 2, 2, 32, 42, 3, 23, 23, 343)
	// _, conclusion := proadder(1, 3, 4, 5, 6, 2, 2, 32, 42, 3, 23, 23, 343)
	fmt.Println("pro result is:", prores)
	fmt.Println(conclusion)
}
