package coder

import "fmt"

func EncodeCoords(latitude string, longitude string) string {
	// binary string is "latitude,longitude"
	binaryString := normalizeCoordinate(latitude)
	binaryString += encodeMap[","]
	binaryString += normalizeCoordinate(longitude)
	fmt.Println(binaryString)

	if result, err := generateWord(binaryString); err != nil {
		fmt.Println(err)
		return ""
	} else {
		return result
	}
}

func Prepare() {
	decodeMap = map[string]string{}
	for key, value := range encodeMap {
		decodeMap[value] = key
	}
}
