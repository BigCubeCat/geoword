package coder

import (
	"fmt"
	"strings"
)

func normalizeCoordinate(coordinate string) string {
	result := ""

	splitString := strings.Split(coordinate, ".")
	beforeDot := splitString[0]
	afterDot := splitString[1]

	// Add prefix for check minus
	if coordinate[0:1] == "-" {
		result += encodeMap["-"]
		beforeDot = beforeDot[1 : len(beforeDot)-1]
	} else {
		result += encodeMap[" "]
	}

	if len(beforeDot) < 3 {
		beforeDot = addPrefixZeros(beforeDot, 3)
	}
	if len(afterDot) < 6 {
		afterDot = addPostfixZeros(afterDot, 6)
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

func generateWord(binaryString string) (string, error) {
	word := ""
	for i := 0; i < len(binaryString)/STEP-1; i++ {
		fmt.Println(
			binaryString[i*STEP : i*STEP+STEP])
		letter, err := binToChar(binaryString[i*STEP : i*STEP+STEP])
		fmt.Println("letter = ", letter)
		if err != nil {
			return "", err
		}
		word += letter
	}
	return word, nil
}
