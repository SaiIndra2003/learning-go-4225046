package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pointers")
	var p *int
	anInt := 42
	p = &anInt
	if p != nil{
		fmt.Println("P is refering to :", *p," P is at ",p)
	} else {
		fmt.Println("P is Nil")
	}

	value1:= 42.3
	pointer1 :=&value1
	fmt.Println("Value 1:",*pointer1, "Located at :",pointer1)
}