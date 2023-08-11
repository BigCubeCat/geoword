package coder

func EncodeCoords(latitude string, longitude string) string {
	// binary string is "latitude,longitude"
	binaryString := normalizeCoordinate(latitude)
	binaryString += encodeMap[","]
	binaryString += normalizeCoordinate(longitude)

	return generateWord(binaryString)
}

func Prepare() {
	decodeMap = map[string]string{}
	for key, value := range encodeMap {
		decodeMap[value] = key
	}
}
