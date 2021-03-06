package pocket

import "encoding/json"
import "time"

// GetQuery configure a retrieve request on Pocket
type GetQuery struct {
	consumerKey string
	accessToken string
	state       QueryState
	favorite    QueryFavorite
	tag         string
	contentType QueryContentType
	sort        QuerySort
	detailType  QueryDetail
	search      string
	domain      string
	since       *time.Time
	count       int
	offset      int
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

// QueryContentType filters on the Pocket items type
type QueryContentType string

// QueryContentType possible values
const (
	QueryContentTypeArticle QueryContentType = "article"
	QueryContentTypeVideo   QueryContentType = "video"
	QueryContentTypeImage   QueryContentType = "image"
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

// QueryDetail level of details of the retrieved Pocket items
type QueryDetail string

// QueryDetail possible values
const (
	QueryDetailSimple   QueryDetail = "simple"
	QueryDetailComplete QueryDetail = "complete"
)

// NewGetQuery initialize a GetQuery
func NewGetQuery(opts ...QueryOption) *GetQuery {
	gq := &GetQuery{
		state:       QueryStateAll,
		favorite:    QueryFavoriteOrNot,
		contentType: QueryContentTypeArticle,
		sort:        QuerySortNewest,
		detailType:  QueryDetailComplete,
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
		DetailType  string `json:"detailType"`
		Search      string `json:"search,omitempty"`
		Domain      string `json:"domain,omitempty"`
		Since       int64  `json:"since,omitempty"`
		Count       int    `json:"count,omitempty"`
		Offset      int    `json:"offset,omitempty"`
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

	j.ContentType = string(gq.contentType)
	j.Sort = string(gq.sort)
	j.DetailType = string(gq.detailType)

	if gq.search != "" {
		j.Search = gq.search
	}

	if gq.domain != "" {
		j.Domain = gq.domain
	}

	if gq.since != nil {
		j.Since = gq.since.Unix()
	}

	if gq.count > 0 {
		j.Count = gq.count
		if gq.offset > 0 {
			j.Offset = gq.offset
		}
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

// WithTag configure the tag filtering
// To retrieve the untagged items use the value _untagged_
func WithTag(tag string) QueryOption {
	return func(gq *GetQuery) {
		gq.tag = tag
	}
}

// WithContentType configure the content type filter
func WithContentType(contentType QueryContentType) QueryOption {
	return func(gq *GetQuery) {
		gq.contentType = contentType
	}
}

// WithSort configure the sort order of the items
func WithSort(sort QuerySort) QueryOption {
	return func(gq *GetQuery) {
		gq.sort = sort
	}
}

// WithDetail configure the level of details returned for each items
func WithDetail(detail QueryDetail) QueryOption {
	return func(gq *GetQuery) {
		gq.detailType = detail
	}
}

// WithSearch to only return items whose title or url contain the search string
func WithSearch(search string) QueryOption {
	return func(gq *GetQuery) {
		gq.search = search
	}
}

// WithDomain to only return items whose title or url contain the search string
func WithDomain(domain string) QueryOption {
	return func(gq *GetQuery) {
		gq.domain = domain
	}
}

// WithSince to only return items modified since the given since unix timestamp
func WithSince(since *time.Time) QueryOption {
	return func(gq *GetQuery) {
		gq.since = since
	}
}

// WithPagination configure number of returned items and an eventual offset
func WithPagination(count, offset int) QueryOption {
	return func(gq *GetQuery) {
		if count < 0 {
			return
		}

		gq.count = count

		if offset > 0 {
			gq.offset = offset
		}
	}
}

func (gq *GetQuery) setCredentials(consumerKey, accessToken string) {
	gq.consumerKey = consumerKey
	gq.accessToken = accessToken
}
