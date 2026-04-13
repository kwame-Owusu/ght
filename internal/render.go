package internal

import (
	"fmt"
	"strings"
	"time"

	"github.com/charmbracelet/lipgloss"
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

	badgePrivateStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("173"))

	cardStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("238")).
			Padding(0, 1).
			MarginBottom(1)

	dividerStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("238"))

	issuesOpenStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("204"))

	commitFreshStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("#6FAF4F"))

	commitAgingStyle = lipgloss.NewStyle().
				Foreground(lipgloss.Color("221"))

	commitDeadStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("204"))

	licenseStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("141"))

	metaStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240")).
			PaddingLeft(1)

	separatorStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("238"))
)

func renderRepos(items []Repository) {
	width := 72
	sep := separatorStyle.Render("│")

	fmt.Println(dividerStyle.Render(strings.Repeat("─", width)))

	for i, item := range items {
		createdAt, _ := time.Parse(time.RFC3339, item.CreatedAt)
		pushedAt, _ := time.Parse(time.RFC3339, item.PushedAt)

		rank := headerStyle.Render(fmt.Sprintf("#%d", i+1))
		name := repoNameStyle.Render(item.FullName)
		badge := renderVisibilityBadge(item.Visibility)
		header := rank + "  " + name + "  " + badge

		desc := descStyle.Render(item.Description)

		stars := starStyle.Render("★ ") + formatNumber(item.StargazersCount)
		openLabel := issuesOpenStyle.Render(fmt.Sprintf("%d open issues", item.OpenIssuesCount))

		commitLine := "Last commit: " + renderCommitAge(pushedAt)

		licenseName := item.LicenseName
		if licenseName == "" {
			licenseName = "None"
		}
		licenseLine := "License: " + licenseStyle.Render(licenseName)

		date := "📅 " + createdAt.Format("Jan 02, 2006")
		url := urlStyle.Render(item.HTMLURL)

		meta := metaStyle.Render(
			lipgloss.JoinVertical(lipgloss.Left,
				stars+"  "+sep+"  "+openLabel,
				date+"  "+sep+"  "+commitLine,
				licenseLine+"  "+sep+"  "+url,
			),
		)

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

func renderVisibilityBadge(v string) string {
	if v == "private" {
		return badgePrivateStyle.Render("[" + v + "]")
	}
	return badgeStyle.Render("[" + v + "]")
}

func renderCommitAge(t time.Time) string {
	days := int(time.Since(t).Hours() / 24)
	var label string
	switch {
	case days == 0:
		label = "today"
	case days == 1:
		label = "yesterday"
	default:
		label = fmt.Sprintf("%d days ago", days)
	}

	switch {
	case days < 30:
		return commitFreshStyle.Render(label)
	case days < 180:
		return commitAgingStyle.Render(label)
	default:
		return commitDeadStyle.Render(label + "  ⚠ possibly abandoned")
	}
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