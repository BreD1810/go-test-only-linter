package stuff

import "fmt"

func UsedFunction() int {
	fmt.Println("This function is used")
	return 1
}

func NotUsedFunction() int {
	fmt.Println("This function is not used")
	return 2
}
