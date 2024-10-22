package main

import (
	"fmt"
)

func main() {
	fmt.Println("using maps:")

	languages := make(map[string]string)
	languages["Python"] = "Python is a high-level, interpreted programming language."
	languages["Java"] = "Java is an object-oriented programming language."
	languages["C++"] = "C++ is a high-performance, compiled programming language."
	languages["JavaScript"] = "JavaScript is a high-level, dynamic programming language."
	fmt.Println(languages) // list of everthing
	// map[C++:C++ is a high-performance, compiled programming language. Java:Java is an object-oriented programming language. JavaScript:JavaScript is a high-level, dynamic programming language. Python:Python is a high-level, interpreted programming language.]

	fmt.Println("js:", languages["JavaScript"])
	delete(languages, "Java") // delete java
	fmt.Println(languages)

	// loops
	//for key Python, value is Python is a high-level, interpreted programming language.
	// for key C++, value is C++ is a high-performance, compiled programming language.
	// for key JavaScript, value is JavaScript is a high-level, dynamic programming language.
	for key, value := range languages { // to ignore key or values we can use _
		fmt.Printf("for key %v, value is %v \n", key, value)
	}
}
