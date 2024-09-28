package gitlab

// ReleaseLink represents a release link.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/releases/links.html
type ReleaseLink struct {
	ID             int           `json:"id"`
	Name           string        `json:"name"`
	URL            string        `json:"url"`
	DirectAssetURL string        `json:"direct_asset_url"`
	External       bool          `json:"external"`
	LinkType       LinkTypeValue `json:"link_type"`
}
