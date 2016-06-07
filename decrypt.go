package main

import (
	"crypto/rc4"
	"encoding/binary"
	"errors"
	"fmt"
	"hash/crc32"
	"os"
	"regexp"
)

type ProtoType uint8

const (
	ProtoWEP ProtoType = iota
	ProtoWPA
	ProtoWPA2
)

type CipherType uint8

const (
	CipherTKIP CipherType = iota
	CipherAES
)

type Data struct {
	Proto  ProtoType
	Cipher CipherType

	Key           []byte
	OriginData    []byte
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

	return decryptWEP(input)
}

func decryptWEP(input *Data) (*Data, error) {
	// switch input.Algorithm {
	// case CipherTKIP:
	// case CipherAES:
	// }

	if len(input.OriginData) < 12 {
		return nil, errors.New("Too short.")
	}

	iv := input.OriginData[:3] // Initialization Vector
	// keyIndex := uint8(d.OriginData[3])     // Key Index
	input.EncryptedData = input.OriginData[4 : len(input.OriginData)-4] // Encrypted data

	keyWithIv := append(iv, input.Key...)

	keyStream, err := rc4.NewCipher(keyWithIv)
	if err != nil {
		return nil, errors.New("Key Stream Create Failed")
	}

	input.DecryptedData = make([]byte, len(input.EncryptedData))
	keyStream.XORKeyStream(input.DecryptedData, input.EncryptedData)

	dec := input.DecryptedData[:len(input.DecryptedData)-4]
	icv := input.DecryptedData[len(input.DecryptedData)-4:]

	h := crc32.NewIEEE()
	_, err = h.Write(dec)
	if err != nil {
		return nil, errors.New("CRC32 Write Failed")
	}

	if h.Sum32() == binary.LittleEndian.Uint32(icv) {
		fmt.Println("!!!")
		fmt.Println(string(dec))
		os.Exit(0)
	} else if h.Sum32() == binary.BigEndian.Uint32(icv) {
		fmt.Println("!!!")
		fmt.Println(string(dec))
		os.Exit(0)
	}

	fmt.Println(string(regexp.MustCompile(`\W`).ReplaceAll(dec, []byte("."))))

	// fmt.Println("iv:", iv, " ki:", ki, " icv:", icv, " fcs:", fcs)
	// fmt.Println("body length:", len(ed))

	return input, nil
}

func decryptWPA(input *Data) (*Data, error) {
	return nil, errors.New("Decrypt Failed")
}

func decryptWPA2(input *Data) (*Data, error) {
	return nil, errors.New("Decrypt Failed")
}
