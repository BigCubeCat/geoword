package coder

import (
	"log"
	"strconv"
	"strings"
)

func encodeBin(latitude string, longitude string) string {
	latitudeBeforeDot, latitudeAfterDot := splitCoordinate(latitude)
	longitudeBeforeDot, longitudeAfterDot := splitCoordinate(longitude)

	binaryString := latitudeBeforeDot[0:1] + longitudeBeforeDot[0:1]

	latitudeBeforeDot = latitudeBeforeDot[1:]
	if len(longitudeBeforeDot) > 3 {
		binaryString += "1"
		longitudeBeforeDot = longitudeBeforeDot[2:]
	} else {
		binaryString += "0"
		longitudeBeforeDot = longitudeBeforeDot[1:]
	}
	number := latitudeBeforeDot + normalizePostfixZeros(latitudeAfterDot, ACCURACY)
	number += longitudeBeforeDot + normalizePostfixZeros(longitudeAfterDot, ACCURACY)
	intNumber, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Println(err)
		return ""
	}
	binaryString += normalizePrefixZeros(strconv.FormatInt(intNumber, 2), BINARYSIZE)
	return binaryString
}

func encodeWord(binaryString string) (string, error) {
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

func decodeWord(word string) string {
	binaryString := ""
	for i := 0; i < len(word); i++ {
		index := int64(strings.Index(ALPHABET, string(word[i])))
		binaryChar := strconv.FormatInt(index, 2)
		binaryChar = normalizePrefixZeros(binaryChar, STEP)
		binaryString += binaryChar
	}
	if binaryString[3:5] != "00" {
		// INVALID word
		log.Println("Invalid word: ", word)
	}
	return binaryString
}

func decodeBinary(binaryString string) (string, string) {
	latitude := decodeSign(binaryString[0:1])
	longitude := decodeSign(binaryString[1:2])
	longitude += binaryString[2:3]

	binString := binaryString[5:] // SKIP empty bits

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
