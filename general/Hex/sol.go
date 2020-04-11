package main

import (
	"encoding/hex"
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

	hexFlag := "63727970746f7b596f755f77696c6c5f62655f776f726b696e675f776974685f6865785f737472696e67735f615f6c6f747d"

	decodedFlag, err := hex.DecodeString(hexFlag)
	checkError(err)

	fmt.Printf("%s\n", decodedFlag)

}