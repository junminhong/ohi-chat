package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

const (
	message = "hello world!"
	secret  = "0933e54e76b24731a2d84b6b463ec04c"
)

func CreatePrivateKey() string {
	/*privateKey := make([]byte, 256)
	if _, err := io.ReadFull(rand.Reader, privateKey); err != nil {
		fmt.Println("Failed to create private key")
	}
	return base64.StdEncoding.EncodeToString(privateKey)*/
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(message))
	//	fmt.Println(h.Sum(nil))
	sha := hex.EncodeToString(h.Sum(nil))
	//	fmt.Println(sha)

	//	hex.EncodeToString(h.Sum(nil))
	return base64.StdEncoding.EncodeToString([]byte(sha))
}
