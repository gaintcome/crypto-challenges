package main

import (
	"encoding/hex"
	"net/http"
	"fmt"
	"os"
	"encoding/json"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func main() {

	// Get ciphertext
	resp, err := http.Get("https://aes.cryptohack.org/block_cipher_starter/encrypt_flag/")
	checkError(err)
	var response map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&response)
	fmt.Println("Ciphertext is = " , response["ciphertext"])


	// Decrypt ciphertext
	resp, err = http.Get("http://aes.cryptohack.org/block_cipher_starter/decrypt/"+response["ciphertext"].(string))
	checkError(err)
	json.NewDecoder(resp.Body).Decode(&response)
	fmt.Println("Plaintext is  = " ,response["plaintext"])


	// Decode hex flag
	decodedFlag, err := hex.DecodeString(response["plaintext"].(string))
	checkError(err)
	fmt.Printf("Flag is = %s\n", decodedFlag)
}
