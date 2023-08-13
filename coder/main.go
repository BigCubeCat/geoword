package coder

import "log"

func EncodeCoords(latitude string, longitude string) string {
	binaryString := encodeBin(latitude, longitude)
	result, err := encodeWord(binaryString)
	if err != nil {
		log.Println(err)
		return ""
	}
	return result
}

func DecodeCoords(word string) (string, string) {
	binaryString := decodeWord(word)
	return decodeBinary(binaryString)
}
