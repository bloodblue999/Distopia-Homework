package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"net/http"
)

const (
	urlPath         = "https://distopia-a1e2.savi2w.workers.dev/"
	customUserAgent = "imperial-penguin"
)

func main() {
	tlsCustomConfig := tls.Config{
		MaxVersion: tls.VersionTLS12,
		MinVersion: tls.VersionTLS12,
	}

	client := &http.Client{
		Transport: &http.Transport{
			DialTLSContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				con, err := tls.Dial(network, addr, &tlsCustomConfig)
				return con, err
			},
		},
	}

	request, errRequest := createRequest()
	if errRequest != nil {
		fmt.Println(errRequest.Error())
		return
	}

	resp, errResponse := client.Do(request)
	if errResponse != nil {
		fmt.Println(errResponse.Error())
		return
	}

	fmt.Println(resp.Status)

	errBodyClose := resp.Body.Close()
	if errBodyClose != nil {
		fmt.Println(errBodyClose.Error())
		return
	}
}

func createRequest() (*http.Request, error) {
	req, err := http.NewRequest(http.MethodGet, urlPath, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", customUserAgent)

	return req, err
}
