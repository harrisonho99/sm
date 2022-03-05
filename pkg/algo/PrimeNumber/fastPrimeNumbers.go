package primenumber

import (
	"math"
)

// check number is prime or not
func CheckPrime(num int) bool {
	result := true

	if num < 2 {
		return false
	}

	sqrt := math.Sqrt(float64(num))
	for i := 2; float64(i) <= sqrt; i++ {
		if num%i == 0 {
			result = false
			break
		}
	}
	return result
}
