package main

import "fmt"

func main() {
	usedFunction()
}

func usedFunction() int {
	fmt.Println("This function is used")
	return 1
}

func notUsedFunction() int {
	fmt.Println("This function is not used")
	return 2
}
