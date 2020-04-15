package main

import (
	"encoding/hex"
	"fmt"
	"os"
	"net"
	"bufio"
	"encoding/json"
	"encoding/base64"
)

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func JSONify(flag []byte)[]byte {

	flagPayload := map[string]string{"decoded": string(flag)}
	jsonPayload, err := json.Marshal(flagPayload)
	checkError(err)
	return jsonPayload

}



type PayloadByte struct{
	Type string
	Encoded []byte
}

type PayloadString struct {
	Type string
	Encoded string
}

func main() {

	conn, err := net.Dial("tcp", "socket.cryptohack.org:13377")
	checkError (err)
	
	connbuf := bufio.NewReader(conn)
	for {
		str, err := connbuf.ReadString('\n')
		checkError(err)

		if len(str)>0 {
			fmt.Print("The challenge is = " , str)
		}

		var data PayloadString	
		json.Unmarshal([]byte(str), &data)

		// Solve base64
		if (data.Type == "base64") {
			decodedFlag, _ := base64.StdEncoding.DecodeString(string(data.Encoded))
			fmt.Printf("Decoded is = %s \n\n", decodedFlag)
			_, err = conn.Write([]byte(string(JSONify(decodedFlag))))
		}

		// Solve hex
		if (data.Type == "hex") {
			decodedFlag, err := hex.DecodeString(data.Encoded)
			checkError(err)
			fmt.Printf("Decoded is = %s \n\n", decodedFlag)	
			_, err = conn.Write([]byte(string(JSONify(decodedFlag))))

		}
		
		// Solve ROT13
		if (data.Type == "rot13") {
			decodedFlag := ""
			for i := 0; i < len(string(data.Encoded)); i++ {
				if ((data.Encoded[i] >= 'A' && data.Encoded[i] <= 'M') || (data.Encoded[i] >= 'a' && data.Encoded[i] <= 'm')) {
					decodedFlag += string(data.Encoded[i]+13)
					continue
				}

				if ((data.Encoded[i] >= 'N' && data.Encoded[i] <= 'Z') || (data.Encoded[i] >= 'n' && data.Encoded[i] <= 'z')) {
					decodedFlag += string(data.Encoded[i]-13)
					continue
				}
				decodedFlag += string(data.Encoded[i])
			}

			fmt.Printf("decoded ROT13 is = %s \n\n", decodedFlag)
			_, err = conn.Write([]byte(string(JSONify([]byte(decodedFlag)))))

		}

		// Solve bigint
		if (data.Type == "bigint") {
			decodedFlag, err := hex.DecodeString(data.Encoded[2:])
			checkError(err)
			fmt.Printf("decoded bigint is = %s \n\n", decodedFlag)	
			_, err = conn.Write([]byte(string(JSONify([]byte(decodedFlag)))))

		}

		// Solve utf-8
		if (data.Type == "utf-8") {
			var data PayloadByte				
			json.Unmarshal([]byte(str), &data)
			fmt.Printf("decoded utf-8 is = %s \n\n", string(data.Encoded))	
			_, err = conn.Write([]byte(string(JSONify((data.Encoded)))))
		}

		checkError(err)
	}

}
