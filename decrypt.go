package main

import (
	"crypto/rc4"
	"encoding/binary"
	"errors"
	"fmt"
)

type ProtoType uint8

const (
	ProtoWEP  ProtoType = 1
	ProtoWPA  ProtoType = 2
	ProtoWPA2 ProtoType = 3
)

type CipherType uint8

const (
	CipherTKIP CipherType = 1
	CipherAES  CipherType = 2
)

type Data struct {
	Proto         ProtoType
	Cipher        CipherType
	Key           []byte
	EncryptedData []byte
	DecryptedData []byte
}

func Decrpyt(input *Data) (*Data, error) {
	// switch input.Proto {
	// case ProtoWEP:
	// 	return decryptWEP(input)
	// case ProtoWPA:
	// 	return decryptWPA(input)
	// case ProtoWPA2:
	// 	return decryptWPA2(input)
	// default:
	// 	if dd, err := decryptWEP(input); err == nil {
	// 		return dd, err
	// 	}
	// 	if dd, err := decryptWPA(input); err == nil {
	// 		return dd, err
	// 	}
	// 	if dd, err := decryptWPA2(input); err == nil {
	// 		return dd, err
	// 	}
	// }

	if data, err := decryptWEP(input); err == nil {
		return data, err
	}

	return nil, errors.New("All decryption attempt fails.")
}

func decryptWEP(input *Data) (*Data, error) {
	// switch input.Algorithm {
	// case CipherTKIP:
	// case CipherAES:
	// }

	iv, _, ed, _, _ := parseWEP(input.EncryptedData)

	keyWithIv := append(iv, input.Key...)
	keyStream, err := rc4.NewCipher(keyWithIv)
	if err != nil {
		return nil, errors.New("Key Stream Create Failed")
	}

	input.DecryptedData = make([]byte, len(ed))
	keyStream.XORKeyStream(input.DecryptedData, ed)

	fmt.Println(string(input.DecryptedData))

	// fmt.Println("iv:", iv, " ki:", ki, " icv:", icv, " fcs:", fcs)
	// fmt.Println("body length:", len(ed))

	return nil, errors.New("Decrypt Failed")
}

func parseWEP(data []byte) ([]byte, []byte, []byte, []byte, uint32) {
	iv := data[:3]                                     // Initialization Vector
	ki := data[3:4]                                    // Key Index
	ed := data[4 : len(data)-8]                        // Encrypted data
	icv := data[len(data)-8 : len(data)-4]             // WEP ICV
	fcs := binary.BigEndian.Uint32(data[len(data)-4:]) // Frame check sequence
	return iv, ki, ed, icv, fcs
}

func decryptWPA(input *Data) (*Data, error) {
	return nil, errors.New("Decrypt Failed")
}

func decryptWPA2(input *Data) (*Data, error) {
	return nil, errors.New("Decrypt Failed")
}
