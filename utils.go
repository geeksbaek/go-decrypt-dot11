package main

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/google/gopacket/pcap"
)

var (
	snaplen = 2048
	timeout = time.Millisecond * 1000
)

func getDeviceFromConsole() pcap.Interface {
	devices, err := pcap.FindAllDevs()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(">> Please select the network card to sniff packets.")
	for i, device := range devices {
		fmt.Printf("\n%d. Name : %s\n   Description : %s\n   IP address : %v\n",
			i+1, device.Name, device.Description, device.Addresses)
	}
	var selected int

	fmt.Print("\n>> ")
	fmt.Scanf("%d", &selected)

	if selected < 0 || selected > len(devices) {
		log.Panic("Invaild Selected.")
	}

	return devices[selected-1]
}

func getPcapHandle(device string) *pcap.Handle {
	inactive, err := pcap.NewInactiveHandle(device)
	if err != nil {
		log.Fatal(err)
	}
	defer inactive.CleanUp()

	if err = inactive.SetRFMon(true); err != nil {
		log.Fatal(err)
	}

	if err = inactive.SetPromisc(true); err != nil {
		log.Fatal(err)
	}

	if err = inactive.SetSnapLen(snaplen); err != nil {
		log.Fatal(err)
	}

	if err = inactive.SetTimeout(timeout); err != nil {
		log.Fatal(err)
	}

	handle, err := inactive.Activate()
	if err != nil {
		log.Fatal(err)
	}

	return handle
}

func dump(_bytes []byte) {
	var b bytes.Buffer
	for i := range _bytes {
		fmt.Fprintf(&b, "%02x ", _bytes[i])
		i := i + 1
		if i != 0 && i%16 == 0 {
			fmt.Fprintf(&b, "\n")
		} else if i != 0 && i%8 == 0 {
			fmt.Fprintf(&b, " ")
		}
	}
	fmt.Println(b.String())
}
