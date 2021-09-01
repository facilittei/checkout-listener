package adapters

import "github.com/go-resty/resty/v2"

// HTTPResty is responsible for making HTTP requests
type HTTPResty struct {
	Client *resty.Client
}

// NewHTTPResty creates a new instance of a HTTPResty adapter
func NewHTTPResty() *HTTPResty {
	return &HTTPResty{
		Client: resty.New(),
	}
}

// Post request
func (requester *HTTPResty) Post(url string, body interface{}, headers map[string]string) ([]byte, error) {
	req := requester.Client.R()
	for key, value := range headers {
		req.SetHeader(key, value)
	}
	req.SetBody(body)
	resp, err := req.Post(url)

	if err != nil {
		return nil, err
	}

	return resp.Body(), nil
}
