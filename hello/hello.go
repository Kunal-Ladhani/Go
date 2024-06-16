package main

// lets the Go compiler know that we want this code to compile and run as a standalone program, as opposed to being a library that's imported by other programs.

import "fmt"

//import fmt imports the fmt (formatting) package. The formatting package exists in Go's standard library and lets us do things like print text to the console.

func main() {
	// func main() defines the main function. main is the name of the function that acts as the entry point for a Go program.
	sayHello("Kunal")
}

func sayHello(name string) {
	fmt.Println("Hello!" + name)
}

// Compilation and Running
// go build -o hello - build go file to executable of custom name
// go build hello.go - directly build go file to executable of same name
// go run .
// go run hello.go - directly run go file
// go run -v hello.go - verbose output
// go run -race hello.go - race detector
