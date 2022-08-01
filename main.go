package main

import (
	"fmt"
)

func main() {
	var name string
	var age uint

	fmt.Println("What's your name?")
	fmt.Scan(&name)
	fmt.Printf("Hello %v", name)
	fmt.Println("What is your age?")
	fmt.Scan(&age)
	fmt.Printf("Nice to meet you %s, I'm also %d years old.", name, age)

	newMessage := fmt.Sprintf("yo %d years is a good age to be, %s", age, name)
	fmt.Println(newMessage)

}
