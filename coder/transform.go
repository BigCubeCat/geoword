package coder

func normalizeCoordinate(coordinate string) string {
	result := ""
	start := 0

	// Add prefix for check minus
	if coordinate[0:1] == "-" {
		result += encodeMap["-"]
		start = 1
	} else {
		result += encodeMap[" "]
	}
	for i := start; i < len(coordinate); i++ {
		result += encodeMap[string(coordinate[i])]
	}
	return result
}
