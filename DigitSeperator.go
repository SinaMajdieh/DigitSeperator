package digitsep



//Seperate digits
func SepDigits(s string) string {
	length := len(s)
	for i := len(s)-1 ; i >= 0 ; i--{
		if s[i] == '.'{
			length = i
			break
		}
	}
	if length < 4 {
		return s
	} else {
		for i := length - 3; i > 0; i -= 3 {
			s = s[:i] + "," + s[i:]
		}
		return s
	}

}
