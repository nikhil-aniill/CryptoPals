package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"log"
)

func main() {
	string1 := []byte("1c0111001f010100061a024b53535009181c")
	string2 := []byte("686974207468652062756c6c277320657965")
	string1Dec, err1 := DecodeHex(string1)
	string2Dec, err2 := DecodeHex(string2)
	if err1 != nil {
		errors.New("Error occurred with string 1")
	}
	if err2 != nil {
		errors.New("Error occurred with string 2")
	}
	decoded, _ := FixedXOR(string1Dec, string2Dec)
	fmt.Printf("The string is %s", EncodeHex(decoded))
}
func FixedXOR(input1 []byte, input2 []byte) ([]byte, error) {
	if len(input1) != len(input2) {
		return nil, errors.New("the inputs have mismatched lengths")
	}
	ret := make([]byte, len(input1))
	for i := 0; i < len(input1); i++ {
		ret[i] = input1[i] ^ input2[i]
	}
	return ret, nil
}

func DecodeHex(input []byte) ([]byte, error) {
	dst := make([]byte, hex.DecodedLen(len(input)))
	_, err := hex.Decode(dst, input) //Decode our string src into our newly created byteslice
	if err != nil {
		log.Fatal("Error occurred")
		return nil, err
	}
	return dst, err
}

func EncodeHex(input []byte) []byte {
	dst := make([]byte, hex.EncodedLen(len(input)))
	hex.Encode(dst, input)
	return dst
}
