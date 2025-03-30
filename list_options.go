package gitlab

import (
	"net/http"
	"net/url"
	"reflect"
	"strconv"
	"strings"
)

type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

const (
	DefaultPerPage = 20
	MaxPerPage     = 100
)

const (
	KeySet = "keyset"
)

type list interface {
	listOptions() ListOptions
}

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
// GitLab API docs: https://docs.gitlab.com/ee/api/rest/index.html#pagination
// https://docs.gitlab.com/ee/api/rest/index.html#pagination-response-headers
// For performance reasons, if a query returns more than 10,000 records, GitLab doesn’t return the following headers:
//
//	x-total
//	x-total-pages
//	rel="last" link
type ListOptions struct {
	// GitLab API docs: https://docs.gitlab.com/ee/api/rest/index.html#offset-based-pagination
	// For paginated result sets, page of results to retrieve.
	Page int `query:"page,omitempty"`
	// For paginated result sets, the number of results to include per page.
	PerPage int `query:"per_page,omitempty"` // default: 20 max: 100

	// GitLab API docs: https://docs.gitlab.com/ee/api/rest/index.html#keyset-based-pagination
	Pagination string             `query:"pagination,omitempty"`
	OrderBy    string             `query:"order_by,omitempty"`
	Sort       Sort               `query:"sort,omitempty"`
	Sets       *map[string]string `query:",inline,omitempty"`
}

func (l ListOptions) listOptions() ListOptions {
	return l
}

func NewListOptions(page int, perPage ...int) ListOptions {
	if page <= 0 {
		page = 1
	}
	pg := DefaultPerPage
	if len(perPage) > 0 && perPage[0] > 0 {
		pg = perPage[0]
	}
	l := ListOptions{
		Page:    page,
		PerPage: pg,
	}

	return l
}

func NewKeySet(orderBy string, sort Sort, perPage ...int) ListOptions {
	pg := DefaultPerPage
	if len(perPage) > 0 && perPage[0] > 0 {
		pg = perPage[0]
	}
	return ListOptions{
		Pagination: KeySet,
		OrderBy:    orderBy,
		Sort:       sort,
		PerPage:    pg,
	}
}

const (
	// Headers used for offset-based pagination.
	xTotal      = "X-Total"
	xTotalPages = "X-Total-Pages"
	xPerPage    = "X-Per-Page"
	xPage       = "X-Page"
	xNextPage   = "X-Next-Page"
	xPrevPage   = "X-Prev-Page"

	// Headers used for keyset-based pagination.
	linkPrev  = "prev"
	linkNext  = "next"
	linkFirst = "first"
	linkLast  = "last"
)

type Page struct {
	ListOptions

	// Fields used for offset-based pagination.
	Total      int
	TotalPages int
	NextPage   int
	PrevPage   int

	// Fields used for keyset-based pagination.
	PrevLink  string
	NextLink  string
	FirstLink string
	LastLink  string
}

func NewPage(l list, resp *http.Response) *Page {
	p := new(Page)
	if reflect.ValueOf(l).IsZero() {
		p.ListOptions = emptyListOptions
	} else {
		p.ListOptions = l.listOptions()
	}

	if resp == nil {
		return p
	}

	p.parse(resp)

	return p
}

func (p *Page) parse(resp *http.Response) {
	if total := resp.Header.Get(xTotal); total != "" {
		if i, err := strconv.Atoi(total); err == nil {
			p.Total = i
		}
	}

	if totalPages := resp.Header.Get(xTotalPages); totalPages != "" {
		if i, err := strconv.Atoi(totalPages); err == nil {
			p.TotalPages = i
		}
	}

	if nextPage := resp.Header.Get(xNextPage); nextPage != "" {
		if i, err := strconv.Atoi(nextPage); err == nil {
			p.NextPage = i
		}
	}

	if prevPage := resp.Header.Get(xPrevPage); prevPage != "" {
		if i, err := strconv.Atoi(prevPage); err == nil {
			p.PrevPage = i
		}
	}

	// keyset
	if p.ListOptions.Pagination == KeySet {
		if link := resp.Header.Get("Link"); link != "" {
			for _, link := range strings.Split(link, ",") {
				parts := strings.Split(link, ";")
				if len(parts) < 2 {
					continue
				}

				linkType := strings.Trim(strings.Split(parts[1], "=")[1], "\"")
				linkValue := strings.Trim(parts[0], "< >")

				switch linkType {
				case linkPrev:
					p.PrevLink = linkValue
				case linkNext:
					p.NextLink = linkValue
				case linkFirst:
					p.FirstLink = linkValue
				case linkLast:
					p.LastLink = linkValue
				}
			}
		}
	}
}

var emptyListOptions = ListOptions{}

func parseLink(link string) url.Values {
	if link == "" {
		return nil
	}
	u, err := url.Parse(link)
	if err != nil {
		return nil
	}
	return u.Query()
}

var keys = []string{"id_after", "cursor"}

func (p *Page) Next() (ListOptions, bool) {
	switch p.ListOptions.Pagination {
	case KeySet:
		values := parseLink(p.NextLink)
		if len(values) == 0 {
			return emptyListOptions, false
		}
		sets := make(map[string]string, len(keys))
		for _, key := range keys {
			if s := values.Get(key); s != "" {
				sets[key] = s
			}
		}
		return ListOptions{
			Pagination: KeySet,
			OrderBy:    p.ListOptions.OrderBy,
			Sort:       p.ListOptions.Sort,
			Sets:       &sets,
		}, true
	default:
		if p.NextPage == 0 {
			return emptyListOptions, false
		}
		return ListOptions{
			Page:    p.NextPage,
			PerPage: p.ListOptions.PerPage,
		}, true
	}
}
