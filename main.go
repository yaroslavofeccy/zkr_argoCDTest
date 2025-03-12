package main

import "fmt"

func getMessage() string {
	return "Hello, Pipe!"
}

func main() {
	msg := getMessage()

	fmt.Println(msg)
	fmt.Println("Successfully!!!")
}
