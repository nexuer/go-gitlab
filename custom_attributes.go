package gitlab

// CustomAttribute struct is used to unmarshal response to api calls.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/custom_attributes.html
type CustomAttribute struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
