package coder

import "strings"

func normalizeCoordinate(coordinate string) string {
	result := ""
	// Add prefix for check minus
	if coordinate[0:1] == "-" {
		result += encodeMap["-"]
	} else {
		result += encodeMap[" "]
	}
	splitString := strings.Split(coordinate, ".")
	beforeDot := splitString[0]
	afterDot := splitString[1]

	if len(beforeDot) < 3 {
		addPrefixZeros(beforeDot, 3)
	}
	if len(afterDot) < 6 {
		addPostfixZeros(afterDot, 6)
	}
	for i := 0; i < len(beforeDot); i++ {
		result += encodeMap[string(beforeDot[i])]
	}
	result += encodeMap["."]
	for i := 0; i < len(afterDot); i++ {
		result += encodeMap[string(afterDot[i])]
	}
	return result
}

func generateWord(binaryString string) string {
	return binaryString
}
