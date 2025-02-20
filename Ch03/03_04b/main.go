package main

import (
	"fmt"
	"sort"
)

func main() {
	// This is an slice
	var colors = []string{"Red", "Green", "Blue"}
	newColors := remove(append(colors, "Black","Pinks"),4)
	sort.Strings(newColors)
	fmt.Println(colors)
	fmt.Println(newColors)
}

func remove(slice []string, i int)[] string{
	return append(slice[:i],slice[i+1:]...)
}
