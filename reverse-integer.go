// tags: math
/** Unique Testcase
 * 1. 1563847412
 * 2. -1563847412
 **/

func reverse(x int) int {
	var result = 0
	var isNegative = x < 0
	if isNegative {
		x *= -1
	}
	for {
		if math.Log2(10.0)+math.Log2(float64(result)) >= 31 {
			result = 0
			break
		}

		result = 10*result + x%10
		x /= 10

		if x == 0 {
			break
		}
	}

	if isNegative {
		return result * -1
	} else {
		return result
	}
}