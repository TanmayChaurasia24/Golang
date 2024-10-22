package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("using slices:")

	var fruitlist = []string{}
	fmt.Printf("type of fruitlist: %T \n", fruitlist) //type of fruitlist: []string

	fruitlist = append(fruitlist, "mango", "banana")
	fmt.Println(fruitlist) // [mango banana]

	fruitlist = append(fruitlist[1:], "apple", "watermelon")
	fmt.Println(fruitlist) // [banana apple watermelon]

	highscore := make([]int, 4)
	fmt.Println(highscore) // [0 0 0 0]

	highscore[0] = 232
	highscore[1] = 323
	highscore[2] = 234
	highscore[3] = 654
	// highscore[4] = 341 // using this will give error out of bound

	highscore = append(highscore, 1232, 4324, 34234) // though array is of size 4 but still append method increases the size of slice
	fmt.Println(highscore)                           // [232 323 234 654 1232 4324 34234]

	sort.Ints(highscore)                       // [232 323 234 654 1232 4324 34234]
	fmt.Println(sort.IntsAreSorted(highscore)) // true

	// how to remove a value from slice based on index
	var course = []string{"reactjs", "ruby", "rust", "c++"}
	fmt.Println(course)                        // [reactjs ruby rust c++]
	course = append(course[:1], course[2:]...) // this will remove the first index element
	fmt.Println(course)                        // [reactjs rust c++]

}
