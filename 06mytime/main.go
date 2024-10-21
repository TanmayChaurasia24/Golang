package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome")

	presenttime := time.Now()

	fmt.Println(presenttime)
	fmt.Println(presenttime.Format("01-02-2006 15:04:05 Monday"))

	createddate := time.Date(2020, time.June, 20, 23, 23, 0, 0, time.Local)
	fmt.Println(createddate.Format("01-02-2006 Monday"))

}

// welcome
// 2024-10-21 20:45:05.2452151 +0530 IST m=+0.000503601
// 10-21-2024 20:45:05 Monday
// 06-20-2020 Saturday
