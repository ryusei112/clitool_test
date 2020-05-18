package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"os"
)

func main() {
	flag.Usage = flagUsage
	encCmd := flag.NewFlagSet("enc", flag.ExitOnError)
	decCmd := flag.NewFlagSet("dec", flag.ExitOnError)

	switch os.Args[1] {
	case "enc":
		s := encCmd.String("s", "", "Encode Base64")
		encCmd.Parse(os.Args[2:])
		sEnc := base64.StdEncoding.EncodeToString([]byte(*s))
		fmt.Println(sEnc)

	case "dec":
		s := decCmd.String("s", "", "Decode Base64")
		decCmd.Parse(os.Args[2:])
		sDec, _ := base64.StdEncoding.DecodeString(*s)
		fmt.Println(string(sDec))

	default:
		flag.Usage()
	}
}

func flagUsage() {
	usageText := `simple cli tool for encode/decode base64.

Usage:
example command [arguments]
The commands are:
enc
dec
Use "base64 [command] --help" for more infomation about a command`

	fmt.Fprintf(os.Stderr, "%s\n\n", usageText)
}
