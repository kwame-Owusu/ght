package internal

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"strings"
	"time"
)

var (
	headerStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("69"))

	repoNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("111"))

	descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")).
			PaddingLeft(1)

	starStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("220"))

	urlStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39"))

	badgeStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("#6FAF4F"))

	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("238")).
			Padding(0, 1).
			MarginBottom(1)

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("238"))
)

func renderRepos(items []Repository) {
	width := 72

	fmt.Println(dividerStyle.Render(strings.Repeat("─", width)))

	for i, item := range items {
		formattedDate, _ := time.Parse(time.RFC3339, item.CreatedAt)

		rank := headerStyle.Render(fmt.Sprintf("#%d", i+1))
		name := repoNameStyle.Render(item.FullName)
		badge := badgeStyle.Render("[" + item.Visibility + "]")

		header := rank + "  " + name + "  " + badge

		desc := descStyle.Render(item.Description)

		stars := starStyle.Render("★ ") + formatNumber(item.StargazersCount)
		date := "  📅 " + formattedDate.Format("Jan 02, 2006")
		url := "  " + urlStyle.Render(item.HTMLURL)

		meta := lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			PaddingLeft(1).
			Render(stars + date + url)

		card := cardStyle.Width(width).Render(
			lipgloss.JoinVertical(lipgloss.Left,
				header,
				"",
				desc,
				meta,
			),
		)

		fmt.Println(card)
	}

	fmt.Println(dividerStyle.Render(
		fmt.Sprintf("── %d results %s", len(items), strings.Repeat("─", width-14)),
	))
}

func formatNumber(n int) string {
	s := fmt.Sprintf("%d", n)
	result := ""
	for i, c := range s {
		if i > 0 && (len(s)-i)%3 == 0 {
			result += ","
		}
		result += string(c)
	}
	return result
}
