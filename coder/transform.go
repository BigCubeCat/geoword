package coder

import (
	"fmt"
	"log"
	"strconv"
)

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
	// splitIndex := 1 + CHAR_SIZE*(3+ACCURACY)

	latitude := decodeSign(binaryString[0:1])
	longitude := decodeSign(binaryString[1:2])
	longitude += binaryString[2:3]
	binString := binaryString[5:] // SKIP empty bits
	fmt.Println(binString, len(binString))
	number, err := strconv.ParseInt(binString, 2, 64)
	if err != nil {
		log.Println(binString, err)
		return "", ""
	}
	lat := strconv.Itoa(int(number / 1000000))
	latitude += lat[0:2] + "." + lat[2:]
	lon := strconv.Itoa(int(number % 1000000))
	longitude += lon[0:2] + "." + lon[2:]

	return latitude, longitude
}
