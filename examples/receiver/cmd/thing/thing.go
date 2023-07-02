package main

import (
	"fmt"
	"test-only-example-receiver/pkg/other"
	"test-only-example-receiver/pkg/stuff"
)

func main() {
	usedFuncMain()
	s := stuff.NewStuffStruct()
	s.UsedFunction()
	s.UsedPointerFunction()
	other.OtherUsedFunc()
	s.UsedFunction()
}

func usedFuncMain() {
	fmt.Println("i am used")
}

func notUsedFuncMain() {
	fmt.Println("im not used")
}
