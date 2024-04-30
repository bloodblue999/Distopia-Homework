package srsamanager

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"hash"
)

type Keys struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	hashFunc   hash.Hash
}

var keys Keys

func GenerateRSAKeyPair() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	publicKey := &privateKey.PublicKey

	keys = Keys{
		publicKey:  publicKey,
		privateKey: privateKey,
		hashFunc:   sha256.New(),
	}
}

func Encrypt(msg string) (string, error) {
	if keys.publicKey == nil {
		return "", errors.New("public key is nil")
	}

	msgBytesEncrypted, err := rsa.EncryptOAEP(keys.hashFunc, rand.Reader, keys.publicKey, []byte(msg), nil)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(msgBytesEncrypted), nil
}

func Decrypt(msg string) (string, error) {
	if keys.privateKey == nil {
		return "", errors.New("private key is nil")
	}

	msgEncryptedBytes, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}

	msgDesencryptedBytes, err := rsa.DecryptOAEP(keys.hashFunc, rand.Reader, keys.privateKey, msgEncryptedBytes, nil)
	if err != nil {
		return "", err
	}

	return string(msgDesencryptedBytes), nil
}
