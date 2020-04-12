package main

import (
	"encoding/hex"
	"encoding/base64"
	"fmt"
	"os"
)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func main() {

	// https://golang.org/pkg/encoding/hex/#Decode
	// https://golang.org/pkg/encoding/base64/

	hexFlag := []byte("72bca9b68fc16ac7beeb8f849dca1d8a783e8acf9679bf9269f7bf")

	// Decoding to bytes
	bytesFlag := make([]byte, hex.DecodedLen(len(hexFlag)))
	_ , err := hex.Decode(bytesFlag, hexFlag)
	checkError(err)

	// Encoding to base64
	flag := make([]byte, base64.StdEncoding.EncodedLen(len(bytesFlag)))
	base64.StdEncoding.Encode(flag, bytesFlag)

	fmt.Printf("%s\n", flag)

}