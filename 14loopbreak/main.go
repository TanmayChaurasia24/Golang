package main

import "fmt"

func main() {
	fmt.Println("using break statement: ")

	days := []string{"sunday", "monday", "friday", "saturday"}
	fmt.Println(days)

	// type 1
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}

	// type 2
	for i := range days {
		fmt.Println(days[i])
	}

	// type 3
	for index, day := range days {
		fmt.Printf("index is %v and value is %v \n", index, day)
	}

	rougevalue := 1

	// type 4
	for rougevalue < 10 {
		if rougevalue == 5 {
			break
		} else if rougevalue == 2 {
			goto lco
		}

		fmt.Println(rougevalue)
		rougevalue++
		// ++rorougevalue // this is wrong
	}

lco:
	fmt.Println("jumped here")
}
