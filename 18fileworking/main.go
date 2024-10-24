package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("working with files in golang")

	content := "this needs to go in a file"

	file, err := os.Create("./myfirstfile.txt") // creating a file
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content) // writing the content in the file
	if err != nil {
		panic(err)
	}

	fmt.Println("length is:", length)
	defer file.Close()
	readfile("./myfirstfile.txt")

}

func readfile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("data inside the file is:", string(databyte))
}
