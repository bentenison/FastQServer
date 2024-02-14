package helpers

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"path"
	"path/filepath"

	"github.com/DataDog/zstd"
)

func SaveDataToFile(filePath string, data []byte, makeDir bool) error {
	if makeDir {
		os.MkdirAll(path.Dir(filePath), os.ModePerm)
	}

	keyBytes := GetKey(filePath)

	compressedText, compressionError := Compress(data)
	if compressionError != nil {
		log.Println("Error while compressing data of ", filePath, " : ", compressionError)
		return compressionError
	}
	encryptedData, encryptionError := Encrypt(compressedText, keyBytes)
	if encryptionError != nil {
		log.Println("Error while encrypting data of ", filePath, " : ", encryptionError)
		return encryptionError
	}
	fileWriteError := WriteFile(filePath, encryptedData)
	if fileWriteError != nil {
		log.Println("Error while writing data to ", filePath, " : ", fileWriteError)
	}
	return fileWriteError
}
func WriteFile(filePath string, data []byte) error {
	return os.WriteFile(filePath, data, 0644)
}

func GetKey(filePath string) []byte {

	defaultSecurityKey := []byte("q{$Kg#+8J(j^x^q8")

	fileName := filepath.Base(filePath)
	fileNameBytes := []byte(fileName)
	fileNameBytes = append(fileNameBytes, defaultSecurityKey...)
	keyBytes := md5.Sum(fileNameBytes)
	return keyBytes[:]
}

func Compress(inputData []byte) ([]byte, error) {
	compressedData, err := zstd.CompressLevel(nil, inputData, 9)
	return compressedData, err
}

func Encrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	padding := block.BlockSize() - len(plainText)%block.BlockSize()

	padtext := bytes.Repeat([]byte{byte(padding)}, padding)

	plainText = append(plainText, padtext...)

	cipherText := make([]byte, len(plainText))

	cbc := cipher.NewCBCEncrypter(block, []byte("TTTTTTTTTTTTTTTT"))
	cbc.CryptBlocks(cipherText, plainText)

	encodedData := make([]byte, base64.StdEncoding.EncodedLen(len(plainText)))

	base64.StdEncoding.Encode(encodedData, cipherText)

	return encodedData, nil
}
func AESDecrypt(encodedData, key []byte) ([]byte, error) {
	decodedData := make([]byte, base64.StdEncoding.DecodedLen(len(encodedData)))
	n, err := base64.StdEncoding.Decode(decodedData, encodedData)
	if err != nil {
		return nil, err
	}
	cipherText := decodedData[:n]

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	cbc := cipher.NewCBCDecrypter(block, []byte("TTTTTTTTTTTTTTTT"))

	if len(cipherText)%cbc.BlockSize() != 0 {
		return nil, errors.New("crypto/cipher: input not full blocks")
	}

	cbc.CryptBlocks(cipherText, cipherText)

	length := len(cipherText)
	if length < 1 {
		log.Println("length of cipherText is less than 1")
		return nil, errors.New("length of cipherText is less than 1")
	}

	unpadding := int(cipherText[length-1])

	if (length - unpadding) < 0 {
		return nil, errors.New("length of (length - unpadding) is less than 0")
	}

	return cipherText[:(length - unpadding)], nil
}
func Decompress(compressedData []byte) ([]byte, error) {
	decompressedData, err := zstd.Decompress(nil, compressedData)
	return decompressedData, err
}
func DecryptData(data []byte, filename string) ([]byte, error) {

	defaultSecurityKey := []byte("q{$Kg#+8J(j^x^q8")

	// securityhelper.SetIV("AAAAAAAAAAAAAAAA")
	fileName := filepath.Base(filename)
	log.Println("Filename", fileName)
	fileNameBytes := []byte(fileName)
	fileNameBytes = append(fileNameBytes, defaultSecurityKey...)
	keyBytes := md5.Sum(fileNameBytes)
	decryptedData, decryptionError := AESDecrypt(data, keyBytes[:])
	if decryptionError != nil {
		log.Println("Error while decrypting data", decryptionError)
		return nil, decryptionError
	}
	decompressedData, decompressError := Decompress(decryptedData)
	if decompressError != nil {
		log.Println("Error while decompressing data", decompressError)
		return nil, decompressError
	}
	return decompressedData, nil
}
