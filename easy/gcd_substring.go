package easy

func gcdOfStrings(str1 string, str2 string) string {
	// if the pattern not repeated, no divisor
	if str1+str2 != str2+str1 {
		return ""
	}

	// get the gcd of the len of str1 and str2
	// so that we can get the length of the str
	// that we can use to divide both strings
	val := gcd(len(str1), len(str2))

	// since the gcd should be dividable by both str
	// simply slice and return one of the strings
	return str1[:val]
}

// calculate the greatest common divisor
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
