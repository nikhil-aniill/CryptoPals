package main

import (
	"encoding/hex"
	"fmt"
)

// try to decode the hex string
func main() {
	hexString, _ := DecodeHex([]byte("1b37373331363f78151b7f2b783431333d78397828372d363c78373e783a393b3736"))
	var answer []byte
	var key int
	var score int
	for i := 0; i < 256; i++ {
		r := make([]byte, len(hexString)) //Converting each integer into byte to XOR with each index of our decoded string
		var s int
		for j := 0; j < len(hexString); j++ {
			b := hexString[j] ^ byte(i)
			s = s + getCharWeight(b)
			r[j] = b
		}
		if s > score {
			key = (i)
			answer = r
			score = s
		}
		s = 0
	}
	fmt.Printf(" Answer is [%s] and the key is [%c]", answer, key)
}
func DecodeHex(input []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dst, input)
	if err != nil {
		return nil, err
	}
	return dst, err
}
func getCharWeight(char byte) int {
	wm := map[byte]int{
		byte('U'): 2,
		byte('u'): 2,
		byte('L'): 3,
		byte('l'): 3,
		byte('D'): 4,
		byte('d'): 4,
		byte('R'): 5,
		byte('r'): 5,
		byte('H'): 6,
		byte('h'): 6,
		byte('S'): 7,
		byte('s'): 7,
		byte(' '): 8,
		byte('N'): 9,
		byte('n'): 9,
		byte('I'): 10,
		byte('i'): 10,
		byte('O'): 11,
		byte('o'): 11,
		byte('A'): 12,
		byte('a'): 12,
		byte('T'): 13,
		byte('t'): 13,
		byte('E'): 14,
		byte('e'): 14,
	}
	return wm[char]
}
