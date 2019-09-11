package digitsep

import (
	"strconv"
)

//Seperate digits
func SepDigits(n int) string {
	s := strconv.Itoa(n)
	if len(s) < 4 {
		return s
	} else {
		for i := len(s) - 3; i >= 0; i -= 3 {
			s = s[:i] + "," + s[i:]
		}
		return s
	}

}
