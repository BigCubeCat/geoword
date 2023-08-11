package coder

import "strconv"

func addPrefixZeros(s string, size int) string {
	for i := 0; i < size-len(s); i++ {
		s = "0" + s
	}
	return s
}

func addPostfixZeros(s string, size int) string {
	for i := 0; i < size-len(s); i++ {
		s += "0"
	}
	return s
}

func binToChar(binaryString string) (string, error) {
	var (
		i   int64
		err error
	)
	if i, err = strconv.ParseInt("1001", 2, 64); err != nil {
		return "0", err
	}
	return string(i), nil
}
