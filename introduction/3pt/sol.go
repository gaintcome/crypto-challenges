package main

import ( 
	"fmt"
	"net"
	"bufio"
    "os"
)


func main() {
	
	// nc socket.cryptohack.org 11111
    conn, err := net.Dial("tcp", "socket.cryptohack.org:11111")
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    
	connbuf := bufio.NewReader(conn)
	for {
	    str, err := connbuf.ReadString('\n')
	    if len(str)>0 {
	        fmt.Print("The flag is = " , str)
	    }
	    if err!= nil {
	        break
		}
	}
}