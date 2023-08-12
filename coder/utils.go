package coder

import "strconv"

func normalizePrefixZeros(s string, size int) string {
	for i := 0; i < size-len(s); i++ {
		s = "0" + s
	}
	if len(s) > size {
		return s[0:size]
	}
	return s
}

func normalizePostfixZeros(s string, size int) string {
	for i := 0; i < size-len(s); i++ {
		s += "0"
	}
	if len(s) > size {
		return s[0:size]
	}
	return s
}

func binToChar(binaryString string) (string, error) {
	if i, err := strconv.ParseInt(binaryString, 2, 64); err != nil {
		return "0", err
	} else {
		index := int(i)
		return string(ALPHABET[index]), nil
	}
}

func getSign(char string) string {
	if char == "1" {
		return "-"
	}
	return ""
}
