package main

import "fmt"

func main() {
	fmt.Println("using array:")

	var fruitlist [4]string

	fruitlist[0] = "apple"
	fruitlist[1] = "banana"
	fruitlist[3] = "orange"

	fmt.Println("fruitlist is:", fruitlist)      // fruitlist is: [apple banana  orange] space for 2 ele
	fmt.Println("fruitlist is:", len(fruitlist)) // 4

	var veglist = [3]string{"potato", "tomato", "onion"}
	fmt.Println("veglist is:", veglist)

}
