/*
A simpler test module
*/
package main

import (
	"fmt"
	"sm/pkg/dslice"
)

func main() {
	sl := dslice.NewSlice()
	dslice.SliceAppend(sl, 30)
	dslice.SliceAppend(sl, 10)
	fmt.Println(sl)
}
