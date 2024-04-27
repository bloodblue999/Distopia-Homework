package main

import (
	"errors"
	"fmt"
	"net/http"
)

func main() {
	urlPath := "https://distopia.savi2w.workers.dev/"

	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			return errors.New("stopping redirect")
		},
	}

	resp, err := client.Get(urlPath)
	if err == nil {
		fmt.Println("Response not redirected")
		return
	}

	if resp == nil {
		fmt.Println("Error in response: ", err.Error())
		return
	}

	fmt.Println(err.Error())
	fmt.Println("Response status: ", resp.Status)
	fmt.Println("Response Distopia header: ", resp.Header.Get("Distopia"))

}
