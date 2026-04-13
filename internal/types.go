package internal

type TrendingParams struct {
	DayFlag   bool
	WeekFlag  bool
	MonthFlag bool
	YearFlag  bool
	Limit     int
	Language  string
}

type SearchResponse struct {
	Items []Repository `json:"items"`
}
 
type Repository struct {
	FullName          string      `json:"full_name"`
	Description       string      `json:"description"`
	HTMLURL           string      `json:"html_url"`
	StargazersCount   int         `json:"stargazers_count"`
	ForksCount        int         `json:"forks_count"`
	Visibility        string      `json:"visibility"`
	CreatedAt         string      `json:"created_at"`
	PushedAt          string      `json:"pushed_at"`
	OpenIssuesCount   int         `json:"open_issues_count"`
	License           *RepoLicense `json:"license"`
	LicenseName       string
}
 
type RepoLicense struct {
	SPDXID string `json:"spdx_id"`
	Name   string `json:"name"`
}