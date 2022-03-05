package gcd

// the largest positive integer that divides each of the integers || highest common factor
func GreatestCommonDivisor(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

// the largest positive integer that divides each of the integers || highest common factor
func GreatestCommonDivisor2(x, y int) int {
	if x <= 0 || y <= 0 {
		return 0
	} else if x == y {
		return x
	} else {
		var result int
		for x > 0 && y > 0 {
			if x == y {
				result = x
				break
			} else if x > y {
				x = x - y
			} else {
				y = y - x
			}
		}
		return result
	}

}
