package encrypt

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"fmt"
)

// Encrypt 对给定字符串加密(AES)
// key 的长度必须为 16|24|32
func Encrypt(plainStr, keyStr string) (cipherStr string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			return
		}
	}()
	keyBytes := []byte(keyStr)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData := PKCS7Padding([]byte(plainStr), blockSize)
	blockMode := cipher.NewCBCEncrypter(block, keyBytes[:blockSize])
	cipherBytes := make([]byte, len(origData))
	blockMode.CryptBlocks(cipherBytes, origData)
	return base64.StdEncoding.EncodeToString(cipherBytes), nil
}

// Encrypt 解密
func Decrypt(cipherStr, keyStr string) (plainStr string, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = errors.New(fmt.Sprint(e))
			return
		}
	}()
	keyBytes := []byte(keyStr)
	cipherBytes, _ := base64.StdEncoding.DecodeString(cipherStr)
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, keyBytes[:blockSize])
	plaintextBytes := make([]byte, len(cipherBytes))
	blockMode.CryptBlocks(plaintextBytes, cipherBytes)
	plaintextBytes = PKCS7UnPadding(plaintextBytes)
	return string(plaintextBytes), nil
}

// PKCS7Padding 补码
func PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// PKCS7UnPadding 去码
func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
