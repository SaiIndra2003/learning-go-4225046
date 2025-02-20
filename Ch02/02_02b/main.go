package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNum := 42
	fmt.Println(str1,str2,str3)
	strLen, err := fmt.Println("The Number is :",aNum)
	if err == nil{
		fmt.Println("String Length is :",strLen)
	}

	fmt.Printf("Type of number : %T\n", aNum)

}
