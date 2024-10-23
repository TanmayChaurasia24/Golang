package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your login count:")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	// Trim the newline character
	input = strings.TrimSpace(input)

	// Convert string input to an integer
	logincount, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input. Please enter a valid number.")
		return
	}

	var result string
	if logincount < 10 {
		result = "regular user"
	} else if logincount >= 10 && logincount < 20 {
		result = "premium user"
	} else {
		result = "admin"
	}

	fmt.Println(result)

	// initializing num and then checking
	if num := 3; num < 10 {
		fmt.Println("num is less than 10")
	} else {
		fmt.Println("num is greater than 10")
	}
}
