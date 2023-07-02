package stuff

import "fmt"

type StuffStruct struct {
	name string
}

func init() {
	fmt.Println("doing an init")
}

func NewStuffStruct() StuffStruct {
	return StuffStruct{name: "testStruct"}
}

func (s StuffStruct) UsedFunction() string {
	fmt.Println("This function is used")
	return s.name
}

func (s *StuffStruct) UsedPointerFunction() string {
	fmt.Println("This function is used")
	return s.name
}

func (s StuffStruct) NotUsedFunction() string {
	fmt.Println("This function is not used")
	return s.name
}

func (s *StuffStruct) NotUsedPointerFunction() string {
	fmt.Println("This function is not used")
	return s.name
}
