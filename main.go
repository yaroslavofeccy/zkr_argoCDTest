package main

import (
	"fmt"
	"time"
)

func getMessage() string {
	return "Hello, Pipe!"
}

func poolMSG(msg string) {
	for msg == "Hello, Pipe!" {
		fmt.Println("Still waiting...")

		time.Sleep(time.Second * 3)
	}
}

func main() {
	msg := getMessage()

	fmt.Println(msg)
	fmt.Println("Successfully!!!")

	fmt.Println("Waiting")
	fmt.Println("Polling!!!")
}
