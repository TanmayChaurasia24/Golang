package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to the ludo game enjoy your game")

	// Seed the random number generator with the current time to ensure different results on each run.
	// This makes the random numbers more unpredictable.
	rand.Seed(time.Now().UnixNano()) // if we comment this out then we will always get 3
	dicenumber := rand.Intn(6) + 1

	switch dicenumber {
	case 1:
		fmt.Println("Dice value is 1 you can open")
	case 2:
		fmt.Println("dice value is 2 moving 2")
	case 3:
		fmt.Println("dice value is 3 moving 3")
	case 4:
		fmt.Println("dice value is 4 moving 4")
	case 5:
		fmt.Println("dice value is 5 moving 5")
	case 6:
		fmt.Println("dice value is 6 moving 6, you are getting another chance")
	default:
		fmt.Println("default case")
	}

}
