package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type TrendingParams struct {
	DayFlag   bool
	WeekFlag  bool
	MonthFlag bool
	YearFlag  bool
	Limit     int
	Language  string
}

func MakeTrendingRequest(p TrendingParams) {
	var duration string
	now := time.Now()

	switch {
	case p.DayFlag:
		duration = now.AddDate(0, 0, -1).Format("2006-01-02")
	case p.WeekFlag:
		duration = now.AddDate(0, 0, -7).Format("2006-01-02")
	case p.MonthFlag:
		duration = now.AddDate(0, -1, 0).Format("2006-01-02")
	case p.YearFlag:
		duration = now.AddDate(-1, 0, 0).Format("2006-01-02")

	}
	query := fmt.Sprintf("created:>%s", duration)
	if p.Language != "" {
		query += "+language:" + p.Language
	}

	url := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=%s&sort=stars&order=desc&per_page=%d",
		query,
		p.Limit,
	)

	searchResp := SearchResponse{}
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	err = json.NewDecoder(resp.Body).Decode(&searchResp)
	if err != nil {
		fmt.Printf("error decoding response into struct: %s", err)
	}

	renderRepos(searchResp.Items)

}
