package main

import (
	"crypto/aes"
	"encoding/hex"
	"fmt"
	"strings"
	"unicode"
)

var (
	q3Key        = "36f18357be4dbd77f050515c73fcf9f2"
	q3Ciphertext = "69dda8455c7dd4254bf353b773304eec0ec7702330098ce7f7520d1cbbb20fc388d1b0adb5054dbd7370849dbf0b88d393f252e764f1f5f7ad97ef79d59ce29f5f51eeca32eabedd9afa9329"
	q4Key        = "36f18357be4dbd77f050515c73fcf9f2"
	q4Ciphertext = "770b80259ec33beb2561358a9f2dc617e46218c0a53cbeca695ae45faa8952aa0e311bde9d4e01726d3184c34451"
)

func main() {
	fmt.Println("----- Question 3 START -----")
	q3CipherHexDecode, err := hexDecode(q3Ciphertext)
	if err != nil {
		panic(err)
	}
	q3KeyHexDecode, err := hexDecode(q3Key)
	if err != nil {
		panic(err)
	}
	q3Decrypted := CTRDecrypt(q3CipherHexDecode, q3KeyHexDecode)
	fmt.Println("Question 3:", outputPrettify(q3Decrypted))
	fmt.Println("----- Question 3 END -----")
	fmt.Println()

	fmt.Println("----- Question 4 START -----")
	q4CipherHexDecode, err := hexDecode(q4Ciphertext)
	if err != nil {
		panic(err)
	}
	q4KeyHexDecode, err := hexDecode(q4Key)
	if err != nil {
		panic(err)
	}
	q4Decrypted := CTRDecrypt(q4CipherHexDecode, q4KeyHexDecode)
	fmt.Println("Question 4:", outputPrettify(q4Decrypted))
	fmt.Println("----- Question 4 END -----")

	/*
		----- Question 3 START -----
		Question 3: CTR mode lets you build a stream cipher from a block cipher.
		----- Question 3 END -----

		----- Question 4 START -----
		Question 4: Always avoid the two time pad!
		----- Question 4 END -----
	*/
}

func hexDecode(input string) ([]byte, error) {
	return hex.DecodeString(input)
}

// outputPrettify avoids printing non-printable characters and gaps
func outputPrettify(input []byte) string {
	var output strings.Builder
	for i := 0; i < len(input); i++ {
		if input[i] < unicode.MaxASCII {
			output.WriteByte(input[i])
		}
	}

	return output.String()
}

func CTRDecrypt(ciphertext, key []byte) []byte {
	// the plaintext will be written here
	var plaintext []byte

	// create a new AES cipher
	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	// get the IV from the ciphertext
	iv := ciphertext[:aes.BlockSize]

	// remove IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// decrypt the ciphertext
	for i := 0; i < len(ciphertext); i += aes.BlockSize {
		// encrypt the IV using the given key
		ivEncrypted := make([]byte, aes.BlockSize)
		block.Encrypt(ivEncrypted, iv)

		// plaintext holder
		plaintextBlock := make([]byte, aes.BlockSize)

		// ciphertext block
		ciphertextBlock := ciphertext[i : i+aes.BlockSize]

		// xor the encrypted IV with the ciphertext
		for j := 0; j < aes.BlockSize; j++ {
			plaintextBlock[j] = ciphertextBlock[j] ^ ivEncrypted[j]
		}

		// add the block to the plaintext
		plaintext = append(plaintext, plaintextBlock...)

		// increase the IV by 1
		for j := len(iv) - 1; j >= 0; j-- {
			iv[j]++
			if iv[j] != 0 {
				break
			}
		}
	}

	return plaintext
}
