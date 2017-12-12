package pocket

import "encoding/json"

// GetQuery configure a retrieve request on Pocket
type GetQuery struct {
	consumerKey string
	accessToken string
	state       QueryState
	favorite    QueryFavorite
	tag         string
	sort        QuerySort
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

// QuerySort sort order of the retrieved Pocket items
type QuerySort string

// QuerySort possible values
const (
	QuerySortNewest QuerySort = "newest"
	QuerySortOldest QuerySort = "oldest"
	QuerySortTitle  QuerySort = "title"
	QuerySortSite   QuerySort = "site"
)

// NewGetQuery initialize a GetQuery
func NewGetQuery(opts ...QueryOption) *GetQuery {
	gq := &GetQuery{
		state:    QueryStateAll,
		favorite: QueryFavoriteOrNot,
		sort:     QuerySortNewest,
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
		Tag         string `json:"tag,omitempty"`
		ContentType string `json:"contentType"`
		Sort        string `json:"sort"`
	}{}
	j.ConsumerKey = gq.consumerKey
	j.AccessToken = gq.accessToken
	j.State = string(gq.state)

	if gq.favorite != QueryFavoriteOrNot {
		favorite := int(gq.favorite)
		j.Favorite = &favorite
	}

	if gq.tag != "" {
		j.Tag = gq.tag
	}

	j.Sort = string(gq.sort)

	// Hardcode ContenType to article as it is the scope of ppocket.
	// Could be changed if video and image should be supported by the ppocket/pocket package.
	j.ContentType = "article"

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

// WithTag configure the tag filtering
// To retrieve the untagged items use the value _untagged_
func WithTag(tag string) QueryOption {
	return func(gq *GetQuery) {
		gq.tag = tag
	}
}

// WithSort configure the sort order of the items
func WithSort(sort QuerySort) QueryOption {
	return func(gq *GetQuery) {
		gq.sort = sort
	}
}

func (gq *GetQuery) setCredentials(consumerKey, accessToken string) {
	gq.consumerKey = consumerKey
	gq.accessToken = accessToken
}
