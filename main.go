package main

import (
	"fmt"
	"time"

	"github.com/google/gopacket"
	"github.com/google/gopacket/pcap"
)

var (
	snaplen = 2048
	timeout = time.Millisecond * 1000
	err     error
	handle  *pcap.Handle
)

func main() {
	device := getDeviceFromConsole()
	handle := getPcapHandle(device.Name)
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		fmt.Println(packet)
	}
}
