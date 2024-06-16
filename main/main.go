package main

// all go files should belong to a package. this one resides in main package

// import the packages that you need - here we are importing fmt - for formatted I/O
import (
	"fmt"
	"main.go/area"
)

// single line comment

/*
	this is multi line comment

	Go == C for 21st century
	2007 by Ken Thompson
	high performance server side applications
	docker, cockroachDB, dGraph etc.
	statically typed complied language
	supports multithreading, concurrency via goroutines

	fast complie time due to advanced dependency analysis
	complies to executable file

	type inference and static typing both supported
	go modules - module system that helps to reuse and export code

*/

/*
	if you define the file name as main.go
	then you can run using - go run .
	it automatically will run the main.go file
	or you can always specify which file you want to run by -
	go run hello.go
	it will directly run the file, and will not create an executable in the folder
	for build you will have to specify the file name -
	go build main.go
	it will create an executable file - main, you can execute it directly -
	./main
*/

// main() function is the entry point for go programs
func main() {

	areaOfCircle := area.Area("circle", 5)
	fmt.Println("area of circle of radius 5 = ", areaOfCircle)

	// declaring variables
	var name string = "kunal" // var variableName datatype = value
	surname := "ladhani"      // go will infer the type automatically
	fmt.Println("-------------------------------------------------")
	fmt.Println("declaring variables")
	fmt.Printf("name is a string, name = %s surname = %s \n", name, surname)

	var num uint32
	num++
	fmt.Printf("this is my num, num = %d \n", num)

	flag := true
	if flag {
		fmt.Printf("boolean val = %t \n", flag)
	}

	var a int
	var b string
	var c bool
	print(a, "\n")
	print(b, "\n")
	print(c, "\n")
	var num1 int32 = 100 // static type
	var num2 = 101       // automatic type inference
	num3 := 102          // automatic type inference - you need to assign a value here
	fmt.Print(num1, num2, num3, "\n")

	fmt.Println("mod op => -40 % 7 = ", (-40 % 7))
}
