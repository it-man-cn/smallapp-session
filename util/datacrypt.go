package util

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

func WXBizDataDecrypt(sessionKey, encryptedData, iv string) (decrypted []byte, err error) {
	var (
		decryptedkey, decryptedIV, decryptedData []byte
	)
	decryptedkey, err = base64.StdEncoding.DecodeString(sessionKey)
	if err != nil {
		return
	}
	decryptedData, err = base64.StdEncoding.DecodeString(encryptedData)
	if err != nil {
		return
	}
	decryptedIV, err = base64.StdEncoding.DecodeString(iv)
	if err != nil {
		return
	}
	block, err := aes.NewCipher(decryptedkey)
	if err != nil {
		fmt.Println(err)
		return
	}
	blockMode := cipher.NewCBCDecrypter(block, decryptedIV)
	decrypted = make([]byte, len(decryptedData))
	// origData := crypted
	blockMode.CryptBlocks(decrypted, decryptedData)
	decrypted = PKCS5UnPadding(decrypted)
	return
}

func WXBizDataSignature(sessionKey, rawData string) string {
	hash := sha1.New()
	_, err := hash.Write([]byte(rawData + sessionKey))
	if err != nil {
		return ""
	}
	sign := hash.Sum(nil)
	return string(sign)
}

// func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
// 	padding := blockSize - len(ciphertext)%blockSize
// 	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
// 	return append(ciphertext, padtext...)
// }

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
