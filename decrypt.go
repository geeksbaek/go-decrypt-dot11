package main

import (
	"errors"
	"fmt"

	"github.com/google/gopacket/layers"
)

func DecryptWEP(dot11WEP layers.Dot11WEP, key string) ([]byte, error) {
	data := dot11WEP.LayerContents()
	iv, ki, ed, icv, fcs := parseWEPBody(data)

	fmt.Println("iv:", iv, " ki:", ki, " icv:", icv, " fcs:", fcs)
	fmt.Println("body length:", len(encryptedData))

	return []byte{}, errors.New("Decrypt Failed")
}

func parseWEPBody(data []byte) ([]byte, []byte, []byte, []byte, []byte) {
	iv := data[:3]                         // Initialization Vector
	ki := data[3:4]                        // Key Index
	ed := data[4 : len(data)-8]            // Encrypted data
	icv := data[len(data)-8 : len(data)-4] // WEP ICV
	fcs := data[len(data)-4:]              // Frame check sequence
	return iv, ki, ed, icv, fcs
}

func DecryptWPA(dot11WEP layers.Dot11WEP, key string) ([]byte, error) {
	return []byte{}, errors.New("Decrypt Failed")
}

func DecryptWPA2(dot11WEP layers.Dot11WEP, key string) ([]byte, error) {
	return []byte{}, errors.New("Decrypt Failed")
}
