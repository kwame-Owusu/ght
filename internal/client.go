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
	Limit     int
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

	}
	url := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=created:>%s&sort=stars&order=desc&per_page=%d",
		duration,
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

	for _, item := range searchResp.Items {
		fmt.Printf("Full Name: %s\n", item.FullName)
		fmt.Printf("Description: %s\n", item.Description)
		fmt.Printf("URL: %s\n", item.HTMLURL)
		fmt.Printf("Creation Date: %s\n", item.CreatedAt)
		fmt.Printf("Visibility: %s\n", item.Visibility)
		fmt.Printf("Stars: %d\n", item.StargazersCount)
		fmt.Println()

	}

}
