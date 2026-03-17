package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func MakeTrendingRequest(dayFlag, weekFlag, monthFlag bool) {
	var duration string
	now := time.Now()

	switch {
	case dayFlag:
		duration = now.AddDate(0, 0, -1).Format("2006-01-02")
	case weekFlag:
		duration = now.AddDate(0, 0, -7).Format("2006-01-02")
	case monthFlag:
		duration = now.AddDate(0, -1, 0).Format("2006-01-02")

	}
	url := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=created:>%s&sort=stars&order=desc&per_page=25",
		duration,
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
		fmt.Println(item.HTMLURL)
	}

}
