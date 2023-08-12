package coder

import (
	"fmt"
)

func EncodeCoords(latitude string, longitude string) string {
	// binary string is "latitude,longitude"
	binaryString := normalizeCoordinate(latitude)
	binaryString += normalizeCoordinate(longitude)
	fmt.Println(binaryString)
	fmt.Println(DecodeCoords(binaryString))
	return binaryString
	/*
		if result, err := generateWord(binaryString); err != nil {
			fmt.Println(err)
			return ""
		} else {
			fmt.Println(DecodeCoords(result))
			return result
		}
	*/
}

func DecodeCoords(word string) string {
	/*
		result := ""
		binaryString := word
			for i := 0; i < len(word); i++ {
				fmt.Println(string(word[i]))
				index := int64(strings.Index(ALPHABET, string(word[i])))
				fmt.Println("index = ", index)
				binaryChar := strconv.FormatInt(index, 2)
				binaryChar = normalizePrefixZeros(binaryChar, STEP)
				fmt.Println("binaryindex = ", binaryChar)
				binaryString += binaryChar
			}
	*/
	return decodeBinary(word)
	/*
		fmt.Println("binaryStrin - ", binaryString)
		for i := 0; i < len(binaryString)/CHAR_SIZE-1; i++ {
			key := binaryString[i*CHAR_SIZE : (i+1)*CHAR_SIZE]
			fmt.Println("key = ", key)
			result += decodeMap[key]
		}
		fmt.Println(result)
		return result

	*/
}

func Prepare() {
	decodeMap = map[string]string{}
	for key, value := range encodeMap {
		decodeMap[value] = key
	}
}
