package main // The main package is the entry point for the Go application

import (
	"bufio"   // Package for buffered I/O operations
	"fmt"     // Package for formatted I/O, similar to printf/scanf in C
	"os"      // Package for OS functions, including input/output operations
	"strconv" // Package for converting between strings and other basic data types
	"strings" // Package for manipulating strings
)

func main() { // The main function where execution of the program begins
	fmt.Println("welcome to the app")           // Print a welcome message to the console
	fmt.Println("enter rating between 1 and 5") // Prompt the user to enter a rating

	reader := bufio.NewReader(os.Stdin) // Create a new buffered reader to read input from standard input (the console)
	input, _ := reader.ReadString('\n') // Read input from the user until a newline character is encountered

	fmt.Println("your rating: ", input)                // Print the user's rating as entered
	fmt.Printf("type of your rating is: %T \n", input) // Print the type of the input variable (should be string)

	// Convert the input string to a float64 number
	numrating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	// Trim whitespace (including the newline) from the input and attempt to convert it to a float64
	if err != nil { // Check if there was an error during conversion
		fmt.Println(err) // Print the error message if conversion fails
	} else { // If no error occurred
		fmt.Println("added 1 to rating: ", numrating+1) // Add 1 to the rating and print the result
	}
}

// $ go run main.go
// welcome to the app
// enter rating between 1 and 5
// 5
// your rating:  5

// type of your rating is: string
// added 1 to rating:  6
