package main

import "fmt"

func PrintBitByteSlice(msg string, b ...byte) {
	fmt.Printf("%-10s: ", msg)
	for i := range b {
		fmt.Printf("%08b ", b[i])
	}
	fmt.Printf("\n")
}

func PrintBitUInt64(msg string, u uint64) {
	fmt.Printf("%-10s: ", msg)

	fmt.Printf("%064b ", u)

	fmt.Printf("\n")
}
