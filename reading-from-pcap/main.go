package main

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	pcapFile = "test.pcap"
	handle   *pcap.Handle
	err      error
)

func main() {
	handle, err = pcap.OpenOffline(pcapFile)
	if err != nil {
		log.Fatal(err)
	}
	defer handle.Close()
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())

}
