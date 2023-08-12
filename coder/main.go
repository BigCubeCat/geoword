package coder

import (
	"fmt"
	"strconv"
	"strings"
)

func EncodeCoords(latitude string, longitude string) string {
	// binary string is "latitude,longitude"
	binaryString := "00"
	binaryString += normalizeCoordinate(latitude)
	binaryString += normalizeCoordinate(longitude)
	fmt.Println(binaryString, len(binaryString))
	if result, err := generateWord(binaryString); err != nil {
		fmt.Println(err)
		return ""
	} else {
		return result
	}
}

func DecodeCoords(word string) (string, string) {
	binaryString := ""
	for i := 0; i < len(word); i++ {
		index := int64(strings.Index(ALPHABET, string(word[i])))
		binaryChar := strconv.FormatInt(index, 2)
		binaryChar = normalizePrefixZeros(binaryChar, STEP)
		fmt.Println("binaryChar = ", binaryChar)
		binaryString += binaryChar
	}
	fmt.Println(binaryString, len(binaryString))
	return decodeBinary(binaryString[SKIP:])
}

func Prepare() {
	decodeMap = map[string]string{}
	for key, value := range encodeMap {
		decodeMap[value] = key
	}
}
