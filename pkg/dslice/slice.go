package dslice

import (
	"errors"
	"fmt"
	"strconv"
)

type ptrToUnderArray interface{}

//untype constant array size -- support size
const (
	size1 = 128
	size2 = 512
	size3 = 1024
	size4 = 2048
	size5 = 4069
	size6 = 8138
)

var errMaxSize = errors.New("slice max size, cannot append anymore")

type slice struct {
	len, cap, elemSize int
	pointer            ptrToUnderArray
}

func SliceFromArray(arrP *[0]int) *slice {
	var p ptrToUnderArray = arrP
	var intSize = 32 // default int in 64bit system
	var sl = slice{len: 0, cap: 0, elemSize: intSize, pointer: p}
	return &sl
}

func NewSlice() *slice {
	var p ptrToUnderArray = [0]int{}
	var intSize = 32 // default int in 64bit system
	var sl = slice{len: 0, cap: 0, elemSize: intSize, pointer: p}
	return &sl
}

func SliceAppend(sl *slice, value int) (*slice, error) {

	// support append to size6 |  8138
	if sl.cap >= size6 {
		return sl, errMaxSize
	}

	//not enough room from array to append
	// malloc new memory
	if sl.cap == sl.len {
		if sl.cap == 0 {
			var newArr [1]int = [1]int{}
			sl.len = 1
			sl.cap = 1
			newArr[0] = value
			sl.pointer = &newArr
			return sl, nil
		} else {
			// use smallest memory range
			// golang need to determine array size at compile time
			//|| size is untype constanst
			return increaseArraySizeAddValue(sl, value)
		}
	} else {
		switch {
		case sl.cap < size1:
			var tempPtr *[1]int = (sl.pointer).(*[1]int)
			sl.len++
			(*tempPtr)[sl.len] = value
			return sl, nil
		case sl.cap < size2:
			var tempPtr *[size1]int = (sl.pointer).(*[size1]int)
			sl.len++
			(*tempPtr)[sl.len] = value
			return sl, nil
		case sl.cap < size3:
			var tempPtr *[size2]int = (sl.pointer).(*[size2]int)
			sl.len++
			(*tempPtr)[sl.len] = value
			return sl, nil
		case sl.cap < size4:
			var tempPtr *[size3]int = (sl.pointer).(*[size3]int)
			sl.len++
			(*tempPtr)[sl.len] = value
			return sl, nil
		case sl.cap < size5:
			var tempPtr *[size4]int = (sl.pointer).(*[size4]int)
			sl.len++
			(*tempPtr)[sl.len] = value
			return sl, nil
		case sl.cap < size6:
			var tempPtr *[size5]int = (sl.pointer).(*[size5]int)
			sl.len++
			(*tempPtr)[sl.len] = value
			return sl, nil
		default:
			return sl, nil
		}

	}
}

func increaseArraySizeAddValue(sl *slice, value int) (*slice, error) {
	switch {
	case sl.cap < size1:
		var newArr [size1]int
		sl.cap = size1
		for i := 0; i < sl.len; i++ {
			var tempPtr *[1]int = (sl.pointer).(*[1]int)
			newArr[i] = (*tempPtr)[i]
		}
		newArr[sl.len] = value
		sl.len++
		sl.pointer = &newArr
		return sl, nil
	case sl.cap < size2:
		var newArr [size2]int
		sl.cap = size2
		for i := 0; i < sl.len; i++ {
			var tempPtr *[size1]int = (sl.pointer).(*[size1]int)
			newArr[i] = (*tempPtr)[i]
		}
		newArr[sl.len] = value
		sl.len++
		sl.pointer = &newArr
		return sl, nil
	case sl.cap < size3:
		var newArr [size3]int
		sl.cap = size3
		for i := 0; i < sl.len; i++ {
			var tempPtr *[size2]int = (sl.pointer).(*[size2]int)
			newArr[i] = (*tempPtr)[i]
		}
		newArr[sl.len] = value
		sl.len++
		sl.pointer = &newArr
		return sl, nil
	case sl.cap < size4:
		var newArr [size4]int
		sl.cap = size4
		for i := 0; i < sl.len; i++ {
			var tempPtr *[size3]int = (sl.pointer).(*[size3]int)
			newArr[i] = (*tempPtr)[i]
		}
		newArr[sl.len] = value
		sl.len++
		sl.pointer = &newArr
		return sl, nil
	case sl.cap < size5:
		var newArr [size5]int
		sl.cap = size5
		for i := 0; i < sl.len; i++ {
			var tempPtr *[size4]int = (sl.pointer).(*[size4]int)
			newArr[i] = (*tempPtr)[i]
		}
		newArr[sl.len] = value
		sl.len++
		sl.pointer = &newArr
		return sl, nil
	case sl.cap < size6:
		var newArr [size6]int
		sl.cap = size6
		for i := 0; i < sl.len; i++ {
			var tempPtr *[size5]int = (sl.pointer).(*[size5]int)
			newArr[i] = (*tempPtr)[i]
		}
		newArr[sl.len] = value
		sl.len++
		sl.pointer = &newArr
		return sl, nil
	default:
		return sl, errMaxSize
	}

}

func (sl *slice) String() string {
	var start string = "[ "
	var end string = " ]"
	var sep string = ", "
	fmt.Println("192----", sl.cap)
	switch {
	case sl.cap < size1:
		var tempPtr *[1]int = (sl.pointer).(*[1]int)
		for i := 0; i < sl.len; i++ {
			if sl.len == i+1 {
				start += strconv.Itoa((*tempPtr)[i]) + end
			} else {
				start += strconv.Itoa((*tempPtr)[i]) + sep
			}
		}
		return start
	case sl.cap < size2:
		var tempPtr *[size1]int = (sl.pointer).(*[size1]int)
		for i := 0; i < sl.len; i++ {
			if sl.len == i+1 {
				start += strconv.Itoa((*tempPtr)[i]) + end
			} else {
				start += strconv.Itoa((*tempPtr)[i]) + sep
			}
		}
		return start
	case sl.cap < size3:
		var tempPtr *[size2]int = (sl.pointer).(*[size2]int)
		for i := 0; i < sl.len; i++ {
			if sl.len == i+1 {
				start += strconv.Itoa((*tempPtr)[i]) + end
			} else {
				start += strconv.Itoa((*tempPtr)[i]) + sep
			}
		}
		return start
	case sl.cap < size4:
		var tempPtr *[size3]int = (sl.pointer).(*[size3]int)
		for i := 0; i < sl.len; i++ {
			if sl.len == i+1 {
				start += strconv.Itoa((*tempPtr)[i]) + end
			} else {
				start += strconv.Itoa((*tempPtr)[i]) + sep
			}
		}
		return start
	case sl.cap < size5:
		var tempPtr *[size4]int = (sl.pointer).(*[size4]int)
		for i := 0; i < sl.len; i++ {
			if sl.len == i+1 {
				start += strconv.Itoa((*tempPtr)[i]) + end
			} else {
				start += strconv.Itoa((*tempPtr)[i]) + sep
			}
		}
		return start
	case sl.cap < size6:
		var tempPtr *[size5]int = (sl.pointer).(*[size5]int)
		for i := 0; i < sl.len; i++ {
			if sl.len == i+1 {
				start += strconv.Itoa((*tempPtr)[i]) + end
			} else {
				start += strconv.Itoa((*tempPtr)[i]) + sep
			}
		}
		return start
	default:
		return start + end
	}

}

//todo
func (sl *slice) GetElement(num int) (element int, ok bool) {
	element, ok = 0, false

	return
}
