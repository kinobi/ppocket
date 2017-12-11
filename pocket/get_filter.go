package pocket

import "encoding/json"

// GetQuery configure a retrieve request on Pocket
type GetQuery struct {
	consumerKey string
	accessToken string
	state       QueryState
	favorite    QueryFavorite
}

// QueryOption give the possibility to configure filters
type QueryOption func(gq *GetQuery)

// QueryState filters on the Pocket item state
type QueryState string

// QueryState possible values
const (
	QueryStateUnread  QueryState = "unread"
	QueryStateArchive QueryState = "archive"
	QueryStateAll     QueryState = "all"
)

// QueryFavorite filters favorite or un-favorite Pocket items
type QueryFavorite int

// QueryFavorite possible values
const (
	QueryFavoriteOrNot    QueryFavorite = -1
	QueryFavoriteExcluded QueryFavorite = 0
	QueryFavoriteOnly     QueryFavorite = 1
)

// NewGetQuery initialize a GetQuery
func NewGetQuery(opts ...QueryOption) *GetQuery {
	gq := &GetQuery{
		state:    QueryStateAll,
		favorite: QueryFavoriteOrNot,
	}

	for _, opt := range opts {
		opt(gq)
	}

	return gq
}

// MarshalJSON marshal GetQuery into valid JSON
func (gq *GetQuery) MarshalJSON() ([]byte, error) {
	j := struct {
		ConsumerKey string `json:"consumer_key"`
		AccessToken string `json:"access_token"`
		State       string `json:"state"`
		Favorite    *int   `json:"favorite,omitempty"`
	}{}
	j.ConsumerKey = gq.consumerKey
	j.AccessToken = gq.accessToken
	j.State = string(gq.state)

	if gq.favorite != QueryFavoriteOrNot {
		favorite := int(gq.favorite)
		j.Favorite = &favorite
	}

	return json.Marshal(j)
}

// WithState configure the state filter
func WithState(state QueryState) QueryOption {
	return func(gq *GetQuery) {
		gq.state = state
	}
}

// WithFavorite configure the favorite filter
func WithFavorite(favorite QueryFavorite) QueryOption {
	return func(gq *GetQuery) {
		gq.favorite = favorite
	}
}

func (gq *GetQuery) setCredentials(consumerKey, accessToken string) {
	gq.consumerKey = consumerKey
	gq.accessToken = accessToken
}
