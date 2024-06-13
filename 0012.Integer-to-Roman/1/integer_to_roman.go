// tags: math
package integer_to_roman

import "strings"

func intToRoman(num int) string {
	var roman string
	units := []string{"I", "V", "X", "L", "C", "D", "M"}

	for i := 0; i <= 3; i++ {
		if num == 0 {
			break
		}
		unit := num % 10
		if unit == 9 {
			roman = units[i*2] + units[i*2+2] + roman
			unit -= 9
		} else if unit == 4 {
			roman = units[i*2] + units[i*2+1] + roman
			unit -= 4
		}

		if unit >= 5 {
			unit -= 5
			roman = units[i*2+1] + strings.Repeat(units[i*2], unit) + roman
		} else {
			roman = strings.Repeat(units[i*2], unit) + roman
		}

		num /= 10
	}

	return roman
}
