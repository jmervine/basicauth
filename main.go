package main

import (
	"flag"
	"fmt"
	"github.com/jmervine/basicauth/basicauth"
	"os"
)

var user, pass string
var header, newline bool

func out(s string) {
	if newline {
		fmt.Println(s)
	} else {
		fmt.Printf("%s", s)
	}
}

func main() {
	flag.StringVar(&user, "u", "", "username")
	flag.StringVar(&pass, "p", "", "password")
	flag.BoolVar(&newline, "n", false, "print newline after header")
	flag.BoolVar(&header, "header", false, "print header string")
	flag.Parse()

	if user == "" || pass == "" {
		fmt.Println("username and password required, see 'help' for details")
		os.Exit(1)
	}

	ba := basicauth.NewBasicAuth(user, pass)

	if header {
		out(ba.Header())
	} else {
		out(ba.String())
	}
}
