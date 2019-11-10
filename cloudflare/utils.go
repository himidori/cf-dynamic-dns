package cloudflare

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	errStatusNotOk = "status code is not 200"
)

func (cr *CFRequester) makeRequest(url, method, authKey, authEmail string, headers map[string]string, data []byte) ([]byte, error) {
	req, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, err
	}

	for k, v := range headers {
		req.Header.Set(k, v)
	}

	req.Header.Set("X-Auth-Email", authEmail)
	req.Header.Set("X-Auth-Key", authKey)

	resp, err := cr.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf(errStatusNotOk)
	}

	return ioutil.ReadAll(resp.Body)
}
