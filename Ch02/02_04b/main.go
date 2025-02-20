package main

import "fmt"

func main() {
	i1, i2, i3 := 12, 45, 68
	intSum := i1 + i2 + i3



	fmt.Println("Total Sum is :", intSum)

	f1, f2, f3:= 23.5,63.1,77.3

	floatSum := f1 + f2 + f3

	fmt.Println("Total Sum is :", floatSum)	

	total := float64(i1) + f1
	
	fmt.Println("Total Sum is :", total)
}
