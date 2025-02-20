package main

import (
	"fmt"
)

func main() {
	fmt.Println(add(1,3))
}

func add(val1 int, val2 int) int{
	return val1 + val2
}