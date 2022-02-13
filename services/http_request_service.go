package services

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

func HttpRequest(endpoint string, object interface{}) error {

	req, err := http.NewRequest(
		http.MethodGet,
		endpoint,
		nil)

	if err != nil {
		return err
	}

	client := &http.Client{Timeout: 5 * time.Second}

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	return json.Unmarshal(body, &object)
}
