package main

import (
	"errors"

	"github.com/google/gopacket/layers"
)

func DecryptWEP(dot11WEP layers.Dot11WEP, key string) ([]byte, error) {
	dump(dot11WEP.LayerContents())

	// data := dot11WEP.LayerContents()

	// data[:4] == Initialization Vector
	// data[4:] == Key Index

	// data[len(data)-8:len(data)-4] == WEP ICV
	// data[len(data)-4:] == Frame check sequence

	return []byte{}, errors.New("Decrypt Failed")
}

func DecryptWPA(dot11WEP layers.Dot11WEP, key string) ([]byte, error) {
	return []byte{}, errors.New("Decrypt Failed")
}

func DecryptWPA2(dot11WEP layers.Dot11WEP, key string) ([]byte, error) {
	return []byte{}, errors.New("Decrypt Failed")
}
