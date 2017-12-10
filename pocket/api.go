package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiRetrieve = "https://getpocket.com/v3/get"

// Item represents a website from the Pocket list
type Item struct {
	URL   string `json:"given_url"`
	Title string `json:"given_title"`
}

// Results represents a result returned by Pocket
type Results struct {
	List map[string]Item
}

// Retrieve execute a call on the retrieve endpoint of Packet
func Retrieve(consumerKey, accessToken string) (*Results, error) {
	input := &map[string]string{
		"consumer_key": consumerKey,
		"access_token": accessToken,
		"count":        "10",
	}

	data, err := json.Marshal(input)
	if err != nil {
		return nil, fmt.Errorf("Failed to encode JSON: %s", err)
	}

	res, err := http.Post(apiRetrieve, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("Failed to call the Pocket API: %s", err)
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response error code %v: %s [%s]", res.StatusCode, res.Header.Get("X-Error"), res.Header.Get("X-Error-code"))
	}

	defer res.Body.Close()
	results := &Results{}
	json.NewDecoder(res.Body).Decode(results)

	return results, nil
}
