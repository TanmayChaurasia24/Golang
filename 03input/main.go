package main // The main package is the entry point for the Go application

import (
	"bufio" // Package for buffered I/O
	"fmt"   // Package for formatted I/O, similar to C's printf/scanf
	"os"    // Package for OS functions, including file and input/output operations
)

func main() { // The main function is where the execution of the program begins
	welcome := "welcome here" // A string variable holding a welcome message
	fmt.Println(welcome)      // Print the welcome message to the console

	reader := bufio.NewReader(os.Stdin) // Create a new buffered reader to read input from standard input (the console)
	fmt.Println("enter the rating: ")   // Prompt the user to enter a rating

	// Read input until a newline character is encountered
	input, _ := reader.ReadString('\n')    // Read a line of input from the user. The second return value (error) is ignored with '_'
	fmt.Println("your rating is: ", input) // Print the user's rating back to the console
}

// $ go run main.go
// welcome here
// enter the rating:
// 10
// your rating is:  10
