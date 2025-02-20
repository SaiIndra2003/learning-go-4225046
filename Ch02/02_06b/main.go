package main

import (
	"fmt"
	"time"
)

func main() {

	t := time.Date(2003, time.May, 18,23, 0, 0, 0, time.UTC)
//	t = time.Now()
	fmt.Printf("I was born at %s\n", t)

	now := time.Now()
	fmt.Printf("Current time %s\n", now)
	fmt.Println("Dates and times")

	fmt.Printf(now.Format(time.ANSIC) + "\n")

}
