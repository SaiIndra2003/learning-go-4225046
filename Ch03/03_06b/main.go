package main

import (
	"fmt"
)

func main() {
	poodle := Dog{"poodle",34}
	bark := Dog{Breed: "Bark",Weight: 33}
	bark.Weight++
	fmt.Println(poodle,bark)
}

type Dog struct{
	Breed string
	Weight int
}
