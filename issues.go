package gitlab

// LabelDetails represents detailed label information.
type LabelDetails struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	Color           string `json:"color"`
	Description     string `json:"description"`
	DescriptionHTML string `json:"description_html"`
	TextColor       string `json:"text_color"`
}

// IssueReferences represents references of the issue.
type IssueReferences struct {
	Short    string `json:"short"`
	Relative string `json:"relative"`
	Full     string `json:"full"`
}
