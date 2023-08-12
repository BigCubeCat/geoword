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
		result += "1"
		beforeDot = beforeDot[1 : len(beforeDot)-1]
	} else {
		result += "0"
	}

	beforeDot = normalizePrefixZeros(beforeDot, 3)
	afterDot = normalizePostfixZeros(afterDot, ACCURACY)
	fmt.Println("before after", beforeDot, afterDot)
	for i := 0; i < len(beforeDot); i++ {
		result += encodeMap[string(beforeDot[i])]
	}
	for i := 0; i < len(afterDot); i++ {
		result += encodeMap[string(afterDot[i])]
	}
	return result
}

func generateWord(binaryString string) (string, error) {
	word := ""
	for i := 0; i < len(binaryString)/STEP; i++ {
		binChar := binaryString[i*STEP : i*STEP+STEP]
		letter, err := binToChar(binChar)
		if err != nil {
			return "", err
		}
		word += letter
	}
	return word, nil
}

func decodeCharByChar(binaryString string) string {
	result := ""
	for i := 0; i < 3; i++ {
		result += decodeMap[binaryString[i*CHAR_SIZE:(i+1)*CHAR_SIZE]]
	}
	result += "."
	for i := 3; i < 3+ACCURACY; i++ {
		result += decodeMap[binaryString[i*CHAR_SIZE:(i+1)*CHAR_SIZE]]
	}
	return result
}

func decodeBinary(binaryString string) (string, string) {
	splitIndex := 1 + CHAR_SIZE*(3+ACCURACY)
	fmt.Println("splitIndex = ", splitIndex)
	fmt.Println(len(binaryString))

	latitude := getSign(binaryString[0:1])
	latitude += decodeCharByChar(binaryString[1:splitIndex])
	longitude := getSign(binaryString[splitIndex : splitIndex+1])
	longitude += decodeCharByChar(binaryString[splitIndex+1:])
	return latitude, longitude
}
