package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("No Ip address argument provided.")
	}
	arg := os.Args[1]
	ip := net.ParseIP(arg)
	if ip == nil {
		log.Fatal("Valid IP not detected, Value provided: " + arg)
	}
	hostnames, err := net.LookupAddr(ip.String())
	if err != nil {
		log.Fatal(err)
	}
	for _, hostname := range hostnames {
		fmt.Println(hostname)
	
	}
}
