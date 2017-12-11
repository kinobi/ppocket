package pocket

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

const apiRetrieve = "https://getpocket.com/v3/get"

// ItemStatus represents the reading state of an item
type ItemStatus string

// ItemStatus possible values
const (
	ItemStatusUnread   ItemStatus = "0"
	ItemStatusArchived ItemStatus = "1"
	ItemStatusDeleted  ItemStatus = "2"
)

// Item represents a website from the Pocket list
type Item struct {
	ItemID        string     `json:"item_id"`
	ResolvedID    string     `json:"resolved_id"`
	GivenURL      string     `json:"given_url"`
	GivenTitle    string     `json:"given_title"`
	Favorite      string     `json:"favorite"`
	Status        ItemStatus `json:"status"`
	ResolvedTitle string     `json:"resolved_title"`
	ResolvedURL   string     `json:"resolved_url"`
	Excerpt       string     `json:"excerpt"`
	IsArticle     string     `json:"is_article"`
	HasVideo      string     `json:"has_video"`
	HasImage      string     `json:"has_image"`
	WordCount     string     `json:"word_count"`
}

// Results represents a result returned by Pocket
type Results struct {
	List map[string]Item
}

// Get execute a call on the retrieve endpoint of Pocket
func Get(consumerKey, accessToken string, gq *GetQuery) (*Results, error) {
	if gq == nil {
		gq = NewGetQuery()
	}

	gq.setCredentials(consumerKey, accessToken)

	data, err := json.Marshal(gq)
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
