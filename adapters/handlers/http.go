package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// HTTP request
type HTTP struct{}

// NewHTTP creates a new instance of the http request
func NewHTTP() *HTTP {
	return &HTTP{}
}

// Post request
func (httpRequest *HTTP) Post(url string, body interface{}, headers map[string]string) ([]byte, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
