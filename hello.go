package main

import "fmt"

//entry point for go, prints Hello, World
// func main() {
// 	fmt.Println("Hello, World!")
// }

/*
	+ %s placeholder is used to represent a string value
	+ Format string requires a string argument to substituted for the %s placeholder
		- the function below is missing that arguement
*/
// func main() {
// 	fmt.Printf("Hello, %s!/n")
// }

func main() {
	fmt.Printf("Hello, %s!/n,", "World")

	fmt.Println("Hello")
	fmt.Println("Greetings and Salutations")
	fmt.Println("Greetings and\n\"Salutations\"")
	fmt.Println(`Greetings and 
	"Salutations"`)

	var smallnumber byte = 120
	fmt.Println(smallnumber)

}

