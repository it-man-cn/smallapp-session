package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
)

var appID = "wx4f4bc4dec97d474b"
var sessionKey = "clnT74VNnN3LyzE99n6rXQ=="
var encryptedData = "Cq83nYbVFh687TaYK5U7DRBNt9j/l70NNeYTh+Yswg5MoTE1TJ21076ltDOOQwQ9X6rse0QwEJm72fpXi9oU/qrHZ42kbrLDgZ5P0HeLZwR2wKVZ3/bIyyFb+PxK/nUGPgjT012sZ3KlMHlCN0jaD/5DV0cJzY7OsEuCXTOWqViJH6zahU0ZP11GOsxSRW/kM3kAOmwt9QgqRb7I69E/zU9YYF/GwODyIjJDJGDtQK43d7GIGuP52JvPVg5Siom8Tip5JzZESN/iCzE91wlz0yLbvOY+O9g7Vgnam46QLH4QDO5eawsh+pj8UPwsfzSEjoc/Y9CX1fwSotXY3BNrR8esXJXxuDIYow+qovJ1AO57nZh8Lcd5WniuQeRUqapN+z7OCj7BBXPLOwgG5BYV4w+BtwZJuFbhDEwVOfkwtTjlEKyMm0kmLK92UX2Nei0vieHJcgOmQtARNuX7BkFra3SkByNhKH1z/aEBzi0VYxwfCsPAmJnP7H/crO6Kyjs3YRsVwUBdNoVpnftz6DUOnA=="
var iv = "fRJ5EFW5dVjHrJzeMiuSRA=="

var rawData = `{"nickName":"陈彬","gender":1,"language":"zh_CN","city":"Minhang","province":"Shanghai","country":"CN","avatarUrl":"http://wx.qlogo.cn/mmopen/vi_32/DYAIOgq83epE9GEPGn7OoOQibQAVs6xly3avcvzYiczL1bWnl4mwalHtY8BmtibQRWmUu9ctKFa86lsK4uXMfaibBA/0"}clnT74VNnN3LyzE99n6rXQ==`

func main() {
	//sessionKey := base64.
	key, _ := base64.StdEncoding.DecodeString(sessionKey)
	data, _ := base64.StdEncoding.DecodeString(encryptedData)
	iv, _ := base64.StdEncoding.DecodeString(iv)
	fmt.Println(key)
	fmt.Println(data)
	fmt.Println(iv)

	block, err := aes.NewCipher(key)
	if err != nil {
		fmt.Println(err)
		return
	}
	//blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(data))
	// origData := crypted
	blockMode.CryptBlocks(origData, data)
	//origData = PKCS5UnPadding(origData)

	fmt.Println(string(origData))

	fmt.Println(rawData)

	hash := sha1.New()
	hash.Write([]byte(rawData))
	bs := hash.Sum(nil)
	fmt.Println(bs)
	fmt.Printf("%x\n", bs)

}
