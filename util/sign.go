package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
)

func GetSha1(message []byte) []byte {
	h := sha1.New()
	h.Write(message)
	return h.Sum(nil)
}

func GenerateSignature(key []byte, iv []byte, message []byte) string {
	messageSha1 := GetSha1(message)

	block, _ := aes.NewCipher(key)

	cbc := cipher.NewCBCEncrypter(block, iv)
	content := PKCS5Padding(messageSha1, block.BlockSize())
	crypted := make([]byte, len(content))
	cbc.CryptBlocks(crypted, content)
	return base64.StdEncoding.EncodeToString(crypted)
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}
