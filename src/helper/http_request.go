package helper

import (
	"io/ioutil"
	"net/http"
)

func HttpRequest(url string, method string, headers map[string]string) []byte {
	client := http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil
	}

	for key, value := range headers {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil
	}

	return body
}
