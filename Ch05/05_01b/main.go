package main

import (
	"fmt"
)

func main() {
	fmt.Println(add("Hi",1,3,33,333))
}

func add( val3 string,vals ...int) int{
	fmt.Println(val3)
	sum := 0
	for _, v := range vals {
		sum += v
	}
	return sum
}