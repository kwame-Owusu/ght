package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

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
 
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
		return
	}
	defer resp.Body.Close()
 
	var searchResp SearchResponse
	if err = json.NewDecoder(resp.Body).Decode(&searchResp); err != nil {
		fmt.Printf("error decoding response into struct: %s\n", err)
		return
	}
 
	for i := range searchResp.Items {
		if searchResp.Items[i].License != nil {
			searchResp.Items[i].LicenseName = searchResp.Items[i].License.SPDXID
		}
	}
 
	renderRepos(searchResp.Items)
}