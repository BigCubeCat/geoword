package coder

var encodeMap = map[string]string{
	"0": "0000",
	"1": "0001",
	"2": "0010",
	"3": "0011",
	"4": "0100",
	"5": "0101",
	"6": "0110",
	"7": "0111",
	"8": "1000",
	"9": "1001",
	".": "1010",
	",": "1011",
	"-": "1100",
	" ": "1101",
}

var decodeMap map[string]string

const ALPHABET = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_."
const SKIP = 2
const STEP = 6
const CHAR_SIZE = 4
const ACCURACY = 4
const BINARYSIZE = 45
