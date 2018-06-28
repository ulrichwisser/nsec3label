package main

import (
	"flag"
	"fmt"
	"os"

	"strconv"

	"github.com/miekg/dns"
)

func main() {

	// define and parse command line arguments
	var verbose bool
	var help bool
	flag.BoolVar(&verbose, "v", false, "print more information while running")
	flag.BoolVar(&help, "h", false, "print usage information")
	flag.Parse()

	if flag.NArg() != 4 || help {
		fmt.Printf("Usage: %s [-v] <hash algorithm number> <iterations> <salt> <fqdn>\n", os.Args[0])
		flag.Usage()
		os.Exit(1)
	}

	var hash uint64
	var iter uint64
	var err error
	var salt string = flag.Arg(2)
	var label string = flag.Arg(3)

	if hash, err = strconv.ParseUint(flag.Arg(0), 10, 8); err != nil {
		panic(err)
	}
	if iter, err = strconv.ParseUint(flag.Arg(1), 10, 16); err != nil {
		panic(err)
	}
	fqdn := dns.Fqdn(label)

	if verbose {
		fmt.Printf("Label for (%s, %d, %d, %s) is %s\n", label, uint8(hash), uint16(iter), salt, dns.HashName(fqdn, uint8(hash), uint16(iter), salt))
	} else {
		fmt.Println(dns.HashName(fqdn, uint8(hash), uint16(iter), salt))
	}
}
