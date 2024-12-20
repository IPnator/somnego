package somnego

import (
	"bytes"
	"crypto/tls"
	"io"
	"net/http"
)

func SendRequest(url string, method string, payload *[]byte) ([]byte, error) {
	tr := &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	client := &http.Client{Transport: tr}
	var req *http.Request

	var bodyData io.Reader
	if payload != nil {
		bodyData = bytes.NewBuffer(*payload)
	}
	req, err := http.NewRequest(method, url, bodyData)
	if err != nil {
		return []byte{}, err
	}

	if method == http.MethodPut || method == http.MethodPost || method == http.MethodPatch {
		req.Header.Set("Content-Type", "application/json")
	}

	// Send request
	res, err := client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return body, nil
}

func GET(url string) ([]byte, error) {
	return SendRequest(url, http.MethodGet, nil)
}
