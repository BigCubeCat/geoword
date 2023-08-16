package coder

import (
	"strconv"
	"strings"
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

func encodeSign(sign string) string {
	if sign == "-" {
		return "1"
	}
	return "0"
}

func decodeSign(char string) string {
	if char == "1" {
		return "-"
	}
	return ""
}

func addSign(coord string) string {
	result := encodeSign(coord[0:1])
	if result == "1" {
		return result + coord[1:]
	}
	return result + coord
}

func splitCoordinate(coordinate string) (string, string) {
	splited := strings.Split(coordinate, ".")
	coordinateBeforeDot := splited[0]
	coordinateAfterDot := splited[1]
	coordinateBeforeDot = addSign(coordinateBeforeDot)
	return coordinateBeforeDot, coordinateAfterDot
}
