package main

import (
	"fmt"
)

func main() {
	fmt.Println("Arrays")
	var colors [3]string 
	colors[0] = "Blue"
	colors[1] = "Pink"
	colors[2] = "Green"
	fmt.Println(colors)

	nums := [5]int{0,1,2,3,4}
	fmt.Println(nums)
}
