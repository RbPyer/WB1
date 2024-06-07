package funcs


func ReverseString(s string) string {
	strRune := []rune(s)
	length := len(strRune)

	for i := 0; i < length / 2; i++ {
		strRune[i], strRune[length - 1 - i] = strRune[length - 1 - i], strRune[i]
	}
	return string(strRune)
}


func ReverseStringV2(s *string) {
	strRune := []rune(*s)
	length := len(strRune)

	for i := 0; i < length / 2; i++ {
		strRune[i], strRune[length - 1 - i] = strRune[length - 1 - i], strRune[i]
	}
	*s = string(strRune)
}
