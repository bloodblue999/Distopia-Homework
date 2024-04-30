package main

import (
	"fmt"
	"github.com/bloodblue999/srsamanager"
)

func main() {
	srsamanager.GenerateRSAKeyPair()
	msg := "MENSAGEM DO SARO"

	encMsg, err := srsamanager.Encrypt(msg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(encMsg)

	decMsg, err := srsamanager.Decrypt(encMsg)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println(decMsg)
}
