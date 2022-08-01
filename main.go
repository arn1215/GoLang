package main

import (
	"fmt"
)

func main() {
	fmt.Println("What's your name?")
	var resp string
	fmt.Scan(&resp)
	fmt.Printf("My name is %v", resp)

}
