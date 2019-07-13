package addition

import (
	"io"
	"net/http"
)

// Request defined Request global proxy
func Request(method string, url string, payload io.Reader) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	return client.Do(req)
}
