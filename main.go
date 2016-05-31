package main

import (
	"fmt"
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

var (
	radioTap      layers.RadioTap
	dot11         layers.Dot11
	dot11WEPLayer layers.Dot11WEP
)

const (
	WEPTestkey = "12345"
	WPATestKey = "qwertyuiop123"
)

// !(wlan.fc.type_subtype == 0x08)

func main() {
	device := getDeviceFromConsole()
	handle := getPcapHandle(device.Name)
	defer handle.Close()

	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		parser := gopacket.NewDecodingLayerParser(
			layers.LayerTypeRadioTap,
			&radioTap,
			&dot11,
			&dot11WEPLayer,
		)
		foundLayerTypes := []gopacket.LayerType{}

		err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
		if err != nil {
			continue
		}

		for _, layerType := range foundLayerTypes {
			switch layerType {

			case layers.LayerTypeRadioTap:
				log.Println("RadioTap Length :", len(radioTap.Contents))
				dump(radioTap.LayerContents())

			case layers.LayerTypeDot11:
				log.Println("Dot11 Length :", len(dot11.Contents))
				dump(dot11.LayerContents())

			case layers.LayerTypeDot11WEP:
				log.Println("Dot11WEP Length :", len(dot11WEPLayer.Contents))
				dump(dot11WEPLayer.LayerContents())

			}
		}

		fmt.Println()
	}
}
