package primenumber

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"strconv"
	"testing"
)

func TestCalculatePrimeNumber(t *testing.T) {

	primeNumMap := fileToDictionary("./prime.txt")

	for i := 0; i < 1224; i++ {
		res := CheckPrime(i)
		_, found := primeNumMap[i]
		if res == found {
			continue
		} else {
			wrongCalulation(i, t)
		}
	}

	fmt.Println("Success calculation")
}

func fileToDictionary(name string) map[int]bool {
	file, err := os.Open(name)
	checkError(err)
	primeNumMap := make(map[int]bool)
	fileScan := bufio.NewScanner(file)

	for fileScan.Scan() {
		num, _ := strconv.Atoi(fileScan.Text())
		primeNumMap[num] = true
	}

	file.Close()

	return primeNumMap
}

func wrongCalulation(i int, t *testing.T) {
	s := fmt.Sprintf("wrong calculation value : %d", i)
	t.Error(s)
}

func checkError(e error) {
	if e != nil {
		debug.PrintStack()
		panic(e.Error())
	}
}
