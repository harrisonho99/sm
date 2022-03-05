package main

import "fmt"

func Switchcase(v int) {

	switch v {
	case Evalcase(1):
		fmt.Println("statement 1")
	case Evalcase(2):
		fmt.Println("statement 2")
	case Evalcase(1):
		fmt.Println("statement 3")
	default:
		fmt.Println("default")
	}
}

/*check case eval or not*/
func Evalcase(num int) int {
	fmt.Printf("eval case %d\n", num)
	return num
}
