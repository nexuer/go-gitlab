package gitlab

type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

// ListOptions specifies the optional parameters to various List methods that
// support pagination.
// GitLab API docs: https://docs.gitlab.com/ee/api/rest/index.html#pagination
type ListOptions struct {
	// GitLab API docs: https://docs.gitlab.com/ee/api/rest/index.html#offset-based-pagination
	// For paginated result sets, page of results to retrieve.
	Page int `query:"page,omitempty" json:"page,omitempty"`
	// For paginated result sets, the number of results to include per page.
	PerPage int `query:"per_page,omitempty" json:"per_page,omitempty"` // default: 20 max: 100

	// GitLab API docs: https://docs.gitlab.com/ee/api/rest/index.html#keyset-based-pagination
	Pagination string `query:"pagination,omitempty" json:"pagination,omitempty"`
	OrderBy    string `query:"order_by,omitempty" json:"order_by,omitempty"`
	Sort       Sort   `query:"sort,omitempty" json:"sort,omitempty"`
}

func NewListOptions(page int, perPage ...int) *ListOptions {
	if page <= 0 {
		page = 1
	}
	l := &ListOptions{
		Page:    page,
		PerPage: 20,
	}
	if len(perPage) > 0 && perPage[0] > 0 {
		l.PerPage = perPage[0]
	}
	return l
}

func NewKeySet(orderBy string, sort Sort, perPage ...int) *ListOptions {
	l := &ListOptions{
		Pagination: "keyset",
		OrderBy:    orderBy,
		Sort:       sort,
	}
	if len(perPage) > 0 && perPage[0] > 0 {
		l.PerPage = perPage[0]
	}
	return l
}
