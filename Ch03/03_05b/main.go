package main

import (
	"fmt"
)

func main() {
	states := make(map[string]string)
	states["Ind"] = "India"
	states["Ch"] = "China"
	fmt.Println(states)
	for k,v := range states{
		fmt.Printf("%v : %v \n",k,v)
	}
	delete(states,"Ch")
	fmt.Println(states)
}
