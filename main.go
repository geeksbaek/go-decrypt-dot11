package main

import (
	"fmt"
	"regexp"
	"time"

	"github.com/google/gopacket"
)

var (
	snaplen = 2048
	timeout = time.Millisecond * 1000
)

func main() {
	device := getDeviceFromConsole()
	handle := getPcapHandle(device.Name)
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Process packet here
		re := regexp.MustCompile(`\W`)
		fmt.Println(re.ReplaceAllString(string(packet.Data()), "."))
	}
}
