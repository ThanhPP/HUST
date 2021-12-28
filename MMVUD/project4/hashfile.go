package main

import (
	"crypto/sha256"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

const (
	blockSize = 1024 // bytes
)

func hash(data []byte) [32]byte {
	return sha256.Sum256(data)
}

func hashFile(name string) (string, error) {
	// open file
	f, err := os.Open(name)
	if err != nil {
		return "", errors.WithMessagef(err, "open file %s error", name)
	}
	defer f.Close()

	// get the file size
	stat, err := f.Stat()
	if err != nil {
		return "", errors.WithMessagef(err, "stat file %s error", name)
	}
	fileSize := int(stat.Size())

	var (
		lastBlockSize = fileSize % blockSize
		lastBlock     = make([]byte, lastBlockSize)
		lastestHash   [32]byte
	)
	// calculate the last block hash
	_, err = f.ReadAt(lastBlock, int64(fileSize-lastBlockSize))
	if err != nil {
		return "", errors.WithMessagef(err, "read the last block error")
	}
	lastestHash = hash(lastBlock)
	fileSize -= lastBlockSize

	// calculate the hash of each block
	var (
		dataBlock = make([]byte, blockSize)
	)
	for fileSize >= blockSize {
		fileSize -= blockSize
		_, err = f.ReadAt(dataBlock, int64(fileSize))
		if err != nil {
			return "", errors.WithMessagef(err, "read block error")
		}
		lastestHash = hash(append(dataBlock, lastestHash[:]...))
	}

	return fmt.Sprintf("%x", lastestHash), nil
}
