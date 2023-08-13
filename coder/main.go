package coder

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func EncodeCoords(latitude string, longitude string) string {
	// binary string is "latitude,longitude"
	lon := strings.Split(longitude, ".")
	fmt.Println("lon = ", lon)
	longitudeBeforeDot := lon[0]
	longitudeAfterDot := lon[1]
	longitudeBeforeDot = addSign(longitudeBeforeDot)

	lat := strings.Split(latitude, ".")
	fmt.Println("lat = ", lat)
	latitudeBeforeDot := lat[0]
	latitudeAfterDot := lat[1]
	latitudeBeforeDot = addSign(latitudeBeforeDot)

	binaryString := latitudeBeforeDot[0:1] + longitudeBeforeDot[0:1]

	latitudeBeforeDot = latitudeBeforeDot[1:]
	if len(longitudeBeforeDot) > 3 {
		binaryString += "1"
		longitudeBeforeDot = longitudeBeforeDot[2:]
	} else {
		binaryString += "0"
		longitudeBeforeDot = longitudeBeforeDot[1:]
	}
	fmt.Println("binaryString = ", binaryString)
	fmt.Println(latitudeBeforeDot, longitudeBeforeDot)
	fmt.Println(len(latitudeBeforeDot), len(longitudeBeforeDot))
	number := latitudeBeforeDot + normalizePostfixZeros(latitudeAfterDot, ACCURACY)
	number += longitudeBeforeDot + normalizePostfixZeros(longitudeAfterDot, ACCURACY)
	intNumber, err := strconv.ParseInt(number, 10, 64)
	if err != nil {
		log.Println(err)
		return ""
	}
	binaryString += normalizePrefixZeros(strconv.FormatInt(intNumber, 2), BINARYSIZE)
	fmt.Println(binaryString, len(binaryString))

	if result, err := generateWord(binaryString); err != nil {
		log.Println(err)
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
		binaryString += binaryChar
	}
	fmt.Println(binaryString, len(binaryString))
	return decodeBinary(binaryString)
}

func Prepare() {
	decodeMap = map[string]string{}
	for key, value := range encodeMap {
		decodeMap[value] = key
	}
}
