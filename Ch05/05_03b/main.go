package main

import (
	"fmt"
	"time"
)

func main() {
	go say("Hello from thread")
	fmt.Println("Hello from main")
	time.Sleep(2 * time.Second)
	
}

func say(message string){
	time.Sleep(1 * time.Second)
	fmt.Println(message)
}
