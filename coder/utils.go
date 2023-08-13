package coder

import (
	"strconv"
)

func normalizePrefixZeros(s string, size int) string {
	result := s
	for i := 0; i < size-len(s); i++ {
		result = "0" + result
	}
	if len(result) > size {
		return result[len(result)-size:]
	}
	return result
}

func normalizePostfixZeros(s string, size int) string {
	result := s
	for i := 0; i < size-len(s); i++ {
		result += "0"
	}
	if len(result) > size {
		return result[0:size]
	}
	return result
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
