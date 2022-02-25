package main

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
	"log"
)

var (
	iface    = "wlo1"
	snaplen  = int32(1600)
	promisc  = false
	timeout  = pcap.BlockForever
	filter   = "tcp and port 80"
	devFound = false
)

func main() {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}
	for _, device := range devices {
		if device.Name == iface {
			devFound = true
		}
	}
	if !devFound {
		log.Fatal("Device not found")
	}
	handle, err := pcap.OpenLive(iface, snaplen, promisc, timeout)
	if err != nil {
		log.Fatal(err)
	}
	source := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range source.Packets() {
		log.Println(packet)
	}
}
