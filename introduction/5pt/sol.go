package main

import ( 
	"net"
	"fmt"
    "os"
    "encoding/json"
    "io/ioutil"

)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}


// nc socket.cryptohack.org 11112
func main() {

    flagPayload := map[string]string{"buy": "flag"}
    jsonPayload, err := json.Marshal(flagPayload)
    // jsonPayload is {"buy":"flag"}
    checkError(err)

    tcpAddr, err := net.ResolveTCPAddr("tcp4", "socket.cryptohack.org:11112")
    checkError(err)

    conn, err := net.DialTCP("tcp", nil, tcpAddr)
    checkError(err)

    // Sending {"buy":"flag"}
    _, err = conn.Write([]byte(jsonPayload))
    checkError(err)


    result, err := ioutil.ReadAll(conn)
    checkError(err)

    fmt.Println(string(result))

}