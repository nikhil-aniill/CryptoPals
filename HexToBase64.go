package main

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	_ "fmt"
	"log"
)

func main() {
	var src []byte
	fmt.Println("Enter Hex Code :- ")
	_, err := fmt.Scanln(&src)
	if err != nil {
		log.Fatal("Input error")
	}
	dst, err := DecodeHex(src)
	if err != nil {
		log.Fatal("Error occurred")
	}
	dstBase64 := EncodeBase64(dst)
	fmt.Printf("The encoded string is :- %s", dstBase64)

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

func EncodeBase64(input []byte) []byte {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(input)))
	base64.StdEncoding.Encode(dst, input)
	return dst
}
