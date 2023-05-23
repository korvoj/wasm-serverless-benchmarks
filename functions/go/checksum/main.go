package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
)

func sha256sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	h := sha256.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func sha512sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	h := sha512.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func md5sum(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	h := md5.New()
	if _, err := io.Copy(h, file); err != nil {
		log.Fatal(err)
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}

func main() {
	if len(os.Args) == 2 {
		filePath := os.Args[1]
		rsha256sum, _ := sha256sum(filePath)
		rsha512sum, _ := sha512sum(filePath)
		rmd5, _ := md5sum(filePath)
		fmt.Printf("\tMD5: %s\t%s\n", rmd5, filePath)
		fmt.Printf("\tSHA256: %s\t%s\n", rsha256sum, filePath)
		fmt.Printf("\tSHA512: %s\t%s\n", rsha512sum, filePath)
	} else {
		log.Fatalln("Please pass a file path as the first argument...")
	}
}
