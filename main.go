package main

import (
	"log"

	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
)

// var (
// 	radioTap     layers.RadioTap
// 	dot11        layers.Dot11
// 	dot11Data    layers.Dot11Data
// 	dot11DataQos layers.Dot11DataQOS
// )

var (
	WEPTestkey = []byte("12345")
	WPATestKey = []byte("qwertyuiop123")
)

// !(wlan.fc.type_subtype == 0x08)

func main() {
	device := getDeviceFromConsole()
	handle := getPcapHandle(device.Name)
	defer handle.Close()

	// Decode a packet
	packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	for packet := range packetSource.Packets() {
		// Get the TCP layer from this packet
		if w := packet.Layer(layers.LayerTypeDot11); w != nil {
			_, err := Decrpyt(&Data{
				OriginData: w.LayerPayload(),
				Key:        WEPTestkey,
				Proto:      ProtoWEP,
			})
			if err != nil {
				log.Println(err)
			}
		}
	}

	// packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
	// for packet := range packetSource.Packets() {
	// 	parser := gopacket.NewDecodingLayerParser(
	// 		layers.LayerTypeRadioTap,
	// 		&radioTap,
	// 		&dot11,
	// 		&dot11Data,
	// 		&dot11DataQos,
	// 	)
	// 	foundLayerTypes := []gopacket.LayerType{}

	// 	err := parser.DecodeLayers(packet.Data(), &foundLayerTypes)
	// 	if err != nil {
	// 		continue
	// 	}

	// 	for _, layerType := range foundLayerTypes {
	// 		switch layerType {

	// 		case layers.LayerTypeRadioTap:
	// 			// log.Println("RadioTap Length :", len(radioTap.Contents))
	// 			// dump(radioTap.LayerContents())

	// 		case layers.LayerTypeDot11:
	// 			// log.Println("Dot11 Length :", len(dot11.Contents))
	// 			// dump(dot11.LayerContents())

	// 		case layers.LayerTypeDot11Data:
	// 			_, err := Decrpyt(&Data{
	// 				OriginData: dot11Data.LayerContents(),
	// 				Key:        WEPTestkey,
	// 				Proto:      ProtoWEP,
	// 			})
	// 			if err != nil {
	// 				log.Println(err)
	// 			}

	// 		case layers.LayerTypeDot11DataQOSData:
	// 			// log.Println("Dot11WEP Length :", len(dot11WEPLayer.Contents))
	// 			// dump(dot11WEPLayer.LayerContents())
	// 			_, err := Decrpyt(&Data{
	// 				OriginData: dot11DataQos.LayerContents(),
	// 				Key:        WEPTestkey,
	// 				Proto:      ProtoWEP,
	// 			})
	// 			if err != nil {
	// 				log.Println(err)
	// 			}
	// 		}
	// 	}
	// }
}
