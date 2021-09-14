package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPRequest is responsible for making HTTP requests
type HTTPRequest struct{}

// NewHTTPRequest creates a new instance of a HTTPRequest adapter
func NewHTTPRequest() *HTTPRequest {
	return &HTTPRequest{}
}

// Post request
func (httpRequest *HTTPRequest) Post(url string, body interface{}, headers map[string]string) ([]byte, error) {
	payload, err := json.Marshal(body)
	if err != nil {
		return nil, handleError(err, "error trying to marshal payload")
	}

	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return nil, handleError(err, "error trying to create an request instance")
	}

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, handleError(err, "error trying to send request")
	}

	defer resp.Body.Close()

	content, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return nil, handleError(err, "error trying read response")
	}

	return content, nil
}

func handleError(err error, message string) error {
	return fmt.Errorf("%v %v", message, err)
}
