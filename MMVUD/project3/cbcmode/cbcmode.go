package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
)

var (
	q1Key        = "140b41b22a29beb4061bda66b6747e14"
	q1Ciphertext = "4ca00ff4c898d61e1edbf1800618fb2828a226d160dad07883d04e008a7897ee2e4b7465d5290d0c0e6c6822236e1daafb94ffe0c5da05d9476be028ad7c1d81"
	q2Key        = "140b41b22a29beb4061bda66b6747e14"
	q2Ciphertext = "5b68629feb8606f9a6667670b75b38a5b4832d0f26e1ab7da33249de7d4afc48e713ac646ace36e872ad5fb8a512428a6e21364b0c374df45503473c5242a253"
)

func main() {
	fmt.Println("----- Question 1 START -----")
	q1CipherHexDecode, err := hexDecode(q1Ciphertext)
	if err != nil {
		panic(err)
	}
	q1KeyHexDecode, err := hexDecode(q1Key)
	if err != nil {
		panic(err)
	}
	q1Decrypted := CBCDecrypt(q1CipherHexDecode, q1KeyHexDecode)
	q1Trimmed := pkcs5Trim(q1Decrypted)
	fmt.Println("Question 1:", string(q1Trimmed))
	fmt.Println("----- Question 1 END -----")
	fmt.Println()

	fmt.Println("----- Question 2 START -----")
	q2CipherHexDecode, err := hexDecode(q2Ciphertext)
	if err != nil {
		panic(err)
	}
	q2KeyHexDecode, err := hexDecode(q2Key)
	if err != nil {
		panic(err)
	}
	q2Decrypted := CBCDecrypt(q2CipherHexDecode, q2KeyHexDecode)
	q2Trimmed := pkcs5Trim(q2Decrypted)
	fmt.Println("Question 2:", string(q2Trimmed))
	fmt.Println("----- Question 2 END -----")

	/*
		----- Question 1 START -----
		Question 1: Basic CBC mode encryption needs padding.
		----- Question 1 END -----

		----- Question 2 START -----
		Question 2: Our implementation uses rand. IV
		----- Question 2 END -----
	*/
}

func hexDecode(hexInput string) ([]byte, error) {
	return hex.DecodeString(hexInput)
}

// pkcs5Trim get the last value of the byte slice and remove the numbers of bytes equal to the last value
func pkcs5Trim(input []byte) []byte {
	padding := input[len(input)-1]
	return input[:len(input)-int(padding)]
}

func printBytes(bytes []byte) {
	for _, b := range bytes {
		fmt.Printf("%08b", b)
	}
	fmt.Println()
}

func CBCDecrypt(ciphertext []byte, key []byte) []byte {
	// the plaintext will be written here
	var plaintext []byte

	// create a new AES cipher with the given key
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// get the initial value (IV)
	// 16-byte encryption IV is chosen at random and is prepended to the ciphertext
	iv := ciphertext[:aes.BlockSize]

	// get the ciphertext (without the IV)
	ciphertext = ciphertext[aes.BlockSize:]

	// each block of the aes cipher is 16 bytes long
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		// get a block from the ciphertext
		ciphertextBlock := ciphertext[i : i+aes.BlockSize]

		// create a plaintext block holder
		var plaintextBlock = make([]byte, aes.BlockSize)

		// decrypt the the ciphertext block and write it to the plaintext block holder
		block.Decrypt(plaintextBlock, ciphertextBlock)

		if i > 0 {
			// get the previous block of the ciphertext
			prevCiphertextBlock := ciphertext[i-aes.BlockSize : i]
			for j := 0; j < aes.BlockSize; j++ {
				// xor the block of decrypted text with the previous block of the ciphertext
				plaintextBlock[j] ^= prevCiphertextBlock[j]
			}
		} else {
			// xor the first block of the decrypted with the IV
			for j := 0; j < aes.BlockSize; j++ {
				plaintextBlock[j] ^= iv[j]
			}
		}

		// append the decrypted block to the plaintext
		plaintext = append(plaintext, plaintextBlock...)
	}

	return plaintext
}
