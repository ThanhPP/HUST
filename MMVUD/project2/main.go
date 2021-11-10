package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
)

func main() {
	var (
		key         = "B4B5B6B7B8B9BABBBCBD" // 10110100 10110101 10110110 10110111 10111000 10111001 10111010 10111011 10111100 10111101
		iv          = "B9BABBBCBD5B6B7B8B4B" // 10111001 10111010 10111011 10111100 10111101 01011011 01101011 01111011 10001011 01001011
		clearFile   = "cleartext.txt"
		cipherFile  = "ciphertext.txt"
		decryptFile = "decrypt.txt"
	)
	// show the key and iv bits
	keyB, err := hex.DecodeString(key)
	if err != nil {
		panic(err)
	}
	ivB, err := hex.DecodeString(iv)
	if err != nil {
		panic(err)
	}

	// encryption
	if err := Encrypt(keyB, ivB, clearFile, cipherFile); err != nil {
		panic(fmt.Errorf("encryption error: %v", err))
	}

	// decryption
	if err := Decrypt(keyB, cipherFile, decryptFile); err != nil {
		panic(fmt.Errorf("decrypt error: %v", err))
	}
}

func Encrypt(key, iv []byte, clearFile, cipherFile string) error {
	// init the trivium
	var (
		inputKey [10]byte
		inputIV  [10]byte
	)
	for i := 0; i < 10; i++ {
		inputKey[i] = key[i]
		inputIV[i] = iv[i]
	}
	trivium := InitTrivium(inputKey, inputIV)

	// create the cipher file
	cipherf, err := os.Create(cipherFile)
	if err != nil {
		return err
	}
	defer cipherf.Close()

	// write the iv into the head of the cleartext
	cipherf.WriteString(hex.EncodeToString(iv))

	// encryption
	clearf, err := os.Open(clearFile)
	if err != nil {
		return err
	}
	defer clearf.Close()

	var (
		buffer = make([]byte, 1)
	)

	for {
		n, err := clearf.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if n == 0 {
			break
		}
		for i := range buffer {
			next := trivium.NextByte()
			buffer[i] ^= next
			hexCipher := hex.EncodeToString(buffer)
			cipherf.WriteString(hexCipher)
		}
	}

	return nil
}

func Decrypt(key []byte, cipherFile, decryptFile string) error {
	// parse the key
	var inputKey [10]byte
	for i := 0; i < 10; i++ {
		inputKey[i] = key[i]
	}

	// open the cipher file
	cipherF, err := os.Open(cipherFile)
	if err != nil {
		return nil
	}
	defer cipherF.Close()

	// get the initial values
	var ivBuf = make([]byte, 20)
	n, err := cipherF.Read(ivBuf)
	if err != nil || n < 20 {
		return fmt.Errorf("Read cipher error: %d - %v", n, err)
	}
	readIV, err := hex.DecodeString(string(ivBuf))
	if err != nil {
		return err
	}
	var inputIV [10]byte
	for i := 0; i < 10; i++ {
		inputIV[i] = readIV[i]
	}

	// init the trivium
	trivium := InitTrivium(inputKey, inputIV)

	// create the decrypt file
	decryptF, err := os.Create(decryptFile)
	if err != nil {
		return err
	}

	// decrypt
	var buf = make([]byte, 4)
	for {
		n, err := cipherF.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if n == 0 {
			break
		}

		decoded, err := hex.DecodeString(string(buf))
		if err != nil {
			return err
		}

		for i := range decoded {
			decrypted := decoded[i] ^ trivium.NextByte()
			decryptF.WriteString(string(decrypted))
		}
	}

	return nil
}
