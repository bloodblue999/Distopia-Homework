package main

import (
	"fmt"
	"github.com/bloodblue999/distopia20240429/srsamanager"
	"io"
	"net/http"
)

func main() {
	createEncryptController()
	createDecryptController()
	srsamanager.GenerateRSAKeyPair()
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func createDecryptController() {
	http.HandleFunc("/decrypt", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.NotFound(w, req)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		msgDecrypted, err := srsamanager.Decrypt(string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprintln(w, msgDecrypted)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Decrypted response body: ", msgDecrypted)
	})
}

func createEncryptController() {
	http.HandleFunc("/encrypt", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != http.MethodPost {
			http.NotFound(w, req)
			return
		}

		body, err := io.ReadAll(req.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		msgEncryptedInBase64, err := srsamanager.Encrypt(string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		_, err = fmt.Fprintln(w, msgEncryptedInBase64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("Encrypted response body: ", msgEncryptedInBase64)
	})
}
