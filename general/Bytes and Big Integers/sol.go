package main

import (
	"fmt"
	"os"
	"math/big"
	"encoding/hex"
)

func checkError(err error) {
    if err != nil {
        fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
        os.Exit(1)
    }
}

func Reverse(s string) string {
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
	    n--
	    runes[n] = rune
	}
	return string(runes[n:])
}

func main() {

	// https://golang.org/pkg/math/big/

	result := ""
	var remainder   = new(big.Int)
	var zero   , _  = new(big.Int).SetString("0", 10)	
	var sixteen, _  = new(big.Int).SetString("16", 10)
	var largenum, _ = new(big.Int).SetString("11515195063862318899931685488813747395775516287289682636499965282714637259206269", 10)	

	// converting base-10 to base-16
	// https://www.permadi.com/tutorial/numDecToHex/
	for {
		remainder.Mod(largenum, sixteen)
		result += fmt.Sprintf("%x",remainder)
		largenum.Div(largenum, sixteen)
		r := largenum.Cmp(zero)
		if r == 0 {
			break
		}
	}

	flag, err := hex.DecodeString(Reverse(result))
	checkError (err)

	fmt.Println("flag is =",string(flag))
}
