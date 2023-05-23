//Adapted from STGDanny - https://gist.github.com/STGDanny/03acf29a90684c2afc9487152324e832
package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	r "math/rand"
	"os"
	"strconv"
	"time"
)

func init() {
	r.Seed(time.Now().UnixNano())
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(length int) string {

	b := make([]byte, length)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

func encrypt(key []byte, message string) (encoded string, err error) {
	//Create byte array from the input string
	plainText := []byte(message)

	//Create a new AES cipher using the key
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//Make the cipher text a byte array of size BlockSize + the length of the message
	cipherText := make([]byte, aes.BlockSize+len(plainText))

	//iv is the ciphertext up to the blocksize (16)
	iv := cipherText[:aes.BlockSize]
	if _, err = io.ReadFull(rand.Reader, iv); err != nil {
		return
	}

	//Encrypt the data:
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(cipherText[aes.BlockSize:], plainText)

	//Return string encoded in base64
	return base64.RawStdEncoding.EncodeToString(cipherText), err
}

func decrypt(key []byte, secure string) (decoded string, err error) {
	//Remove base64 encoding:
	cipherText, err := base64.RawStdEncoding.DecodeString(secure)

	//IF DecodeString failed, exit:
	if err != nil {
		return
	}

	//Create a new AES cipher with the key and encrypted message
	block, err := aes.NewCipher(key)

	//IF NewCipher failed, exit:
	if err != nil {
		return
	}

	//IF the length of the cipherText is less than 16 Bytes:
	if len(cipherText) < aes.BlockSize {
		err = errors.New("Ciphertext block size is too short!")
		return
	}

	iv := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]

	//Decrypt the message
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(cipherText, cipherText)

	return string(cipherText), err
}

func main() {
	if len(os.Args) == 3 {
		randomStringLength, _ := strconv.Atoi(os.Args[1])
		message := randomString(randomStringLength)
		iterations, _ := strconv.Atoi(os.Args[2])
		fmt.Printf("\tCLEAR TEXT MESSAGE:%s\n\n\n", message)
		for i := 0; i < iterations; i++ {
			cipherKey := []byte("\xa1\xf6%\x8c\x87}_\xcd\x89dHE8\xbf\xc9,") //16 bit key for AES-128
			encrypted, err := encrypt(cipherKey, message)
			//IF the encryption failed:
			if err != nil {
				//Print error message:
				log.Println(err)
				os.Exit(-2)
			}

			//Print the key and cipher text:
			//fmt.Printf("\n\tCIPHER KEY: %s\n", string(cipherKey))
			fmt.Printf("\tENCRYPTED: %s\n", encrypted)

			//Decrypt the text:
			decrypted, err := decrypt(cipherKey, encrypted)

			//IF the decryption failed:
			if err != nil {
				log.Println(err)
				os.Exit(-3)
			}

			//Print re-decrypted text:
			fmt.Printf("\tDECRYPTED: %s\n\n", decrypted)
		}
	} else {
		log.Fatalln("Exactly two arguments are required: <string_length> <iterations>...")
	}
}
