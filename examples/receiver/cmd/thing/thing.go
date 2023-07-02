package main

import (
	"test-only-example-receiver/pkg/other"
	"test-only-example-receiver/pkg/stuff"
)

func main() {
	s := stuff.NewStuffStruct()
	s.UsedFunction()
	s.UsedPointerFunction()
	other.OtherUsedFunc()
	s.UsedFunction()
}
