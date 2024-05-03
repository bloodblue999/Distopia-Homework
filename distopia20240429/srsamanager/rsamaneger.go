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

type keys struct {
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	hashFunc   hash.Hash
}

var keysValue keys

func GenerateRSAKeyPair() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	publicKey := &privateKey.PublicKey

	keysValue = keys{
		publicKey:  publicKey,
		privateKey: privateKey,
		hashFunc:   sha256.New(),
	}
}

func Encrypt(msg string) (string, error) {
	if keysValue.publicKey == nil {
		return "", errors.New("public key is nil")
	}

	msgBytesEncrypted, err := rsa.EncryptOAEP(keysValue.hashFunc, rand.Reader, keysValue.publicKey, []byte(msg), nil)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(msgBytesEncrypted), nil
}

func Decrypt(msg string) (string, error) {
	if keysValue.privateKey == nil {
		return "", errors.New("private key is nil")
	}

	msgEncryptedBytes, err := base64.StdEncoding.DecodeString(msg)
	if err != nil {
		return "", err
	}

	msgDecryptedBytes, err := rsa.DecryptOAEP(keysValue.hashFunc, rand.Reader, keysValue.privateKey, msgEncryptedBytes, nil)
	if err != nil {
		return "", err
	}

	return string(msgDecryptedBytes), nil
}
