package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "welcome here"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("enter the rating: ")

	// comma error syntax
	input, _ := reader.ReadString('\n')
	fmt.Println("your rating is: ", input)
}

// $ go run main.go
// welcome here
// enter the rating:
// 10
// your rating is:  10
