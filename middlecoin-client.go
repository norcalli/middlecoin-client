package main

import (
	"flag"
	"fmt"
	"github.com/norcalli/middlecoin"
	"log"
)

const url = "http://www.middlecoin.com/json"

func main() {
	flag.Parse()
	addresses := flag.Args()
	if len(addresses) == 0 {
		fmt.Println("Usage: middlecoin-client address1 address2 ...")
		return
	}

	r, err := middlecoin.Fetch()
	if err != nil {
		log.Fatal(err)
	}
	total := new(middlecoin.AddressReport)
	for _, address := range addresses {
		fmt.Println(address)
		total.Add(r.Report[address])
		fmt.Println(r.Report[address].String())
	}
	if len(addresses) > 1 {
		fmt.Println(total.String())
	}
}
