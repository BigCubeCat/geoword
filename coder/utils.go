package coder

import "strconv"

func addPrefixZeros(s string, size int) string {
	for i := 0; i < size-len(s); i++ {
		s = "0" + s
	}
	return s
}

func addPostfixZeros(s string, size int) string {
	for i := 0; i < size-len(s); i++ {
		s += "0"
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
