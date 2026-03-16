package internal

type Repository struct {
	FullName        string `json:"full_name"`
	HTMLURL         string `json:"html_url"`
	Description     string `json:"description"`
	CreatedAt       string `json:"created_at"`
	StargazersCount int    `json:"stargazers_count"`
	Visibility      string `json:"visibility"`
}

type SearchResponse struct {
	TotalCount        int          `json:"total_count"`
	IncompleteResults bool         `json:"incomplete_results"`
	Items             []Repository `json:"items"`
}
