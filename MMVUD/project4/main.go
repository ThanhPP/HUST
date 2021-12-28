package main

import (
	"fmt"
	"log"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "project4",
		Short: "project4 is a tool to calculate the hash of a file or compare file and hashes",
	}
)

func init() {
	rootCmd.AddCommand(calculateHashCmd)
	rootCmd.AddCommand(compareFilesCmd)
	compareHashCmd.PersistentFlags().StringVarP(&expectedHash, "expected", "e", "", "expected hash")
	rootCmd.AddCommand(compareHashCmd)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		return
	}
}

var calculateHashCmd = &cobra.Command{
	Use:   "calhash",
	Short: "calculate the hash of files",
	RunE:  calculateHash,
}

func calculateHash(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		errors.New("empty file name")
	}

	for i := range args {
		log.Println("file name:", args[i])
		fileHash, err := hashFile(args[i])
		if err != nil {
			return err
		}
		log.Println("hash:", fileHash)
	}

	return nil
}

var compareFilesCmd = &cobra.Command{
	Use:   "comparefiles",
	Short: "compare the hash of 2 files",
	Long:  "the arguments of this command need exactly 2 files",
	RunE:  compareFiles,
}

func compareFiles(cmd *cobra.Command, args []string) error {
	if len(args) != 2 {
		return errors.New("invaid arguments length (need exactly 2 files)")
	}

	log.Println("file name:", args[0])
	file1Hash, err := hashFile(args[0])
	if err != nil {
		return errors.WithMessagef(err, "hash file %s error", args[0])
	}
	log.Println("hash:", file1Hash)

	log.Println("file name:", args[1])
	file2Hash, err := hashFile(args[1])
	if err != nil {
		return errors.WithMessagef(err, "hash file %s error", args[0])
	}
	log.Println("hash:", file2Hash)

	log.Println("file1 hash == file2 hash:", file1Hash == file2Hash)

	return nil
}

var (
	expectedHash   string
	compareHashCmd = &cobra.Command{
		Use:   "comparehash",
		Short: "compare the hash of a file with an expected hash",
		Long:  "the arguments of this command need exactly 1 file and 1 flag is the expected hash",
		RunE:  compareHash,
	}
)

func compareHash(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("invaid arguments length (need exactly 1 file)")
	}

	if len(expectedHash) != 64 {
		return fmt.Errorf("expected hash length %d != %d", len(expectedHash), 64)
	}

	log.Println("file name:", args[0])
	fileHash, err := hashFile(args[0])
	if err != nil {
		return errors.WithMessagef(err, "hash file %s error", args[0])
	}
	log.Println("hash:         ", fileHash)
	log.Println("expected hash:", expectedHash)

	log.Println("file hash == expected hash:", fileHash == expectedHash)

	return nil
}
