package internal

import (
	"fmt"
	"net/http"
	"time"
)

func MakeDayRequest() {
	now := time.Now()
	day := now.AddDate(0, 0, -1).Format("2006-01-02")
	// week := now.AddDate(0, 0, -7).Format("2006-01-02")
	// month := now.AddDate(0, -1, 0).Format("2006-01-02")

	url := fmt.Sprintf(
		"https://api.github.com/search/repositories?q=created:>%s&sort=stars&order=desc&per_page=25",
		day, // swap for day or month
	)
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error making http request: %s\n", err)
	}

	fmt.Println(res.Body)
}
