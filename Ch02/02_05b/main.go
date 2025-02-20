package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	sum = math.Round(sum*100) / 100
	fmt.Println("Float sum:", sum)

}
