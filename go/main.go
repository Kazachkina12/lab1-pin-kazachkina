package main

import "fmt"

// greet returns a greeting for the specified name.
func greet(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

// main is the entry point of the program.
func main() {
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)

	message := greet(userName)
	fmt.Println(message)
}
