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
			Foreground(lipgloss.Color("69")). // soft purple/blue
			PaddingLeft(1)

	repoNameStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("111")) // light blue

	descStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("245")). // muted gray
			PaddingLeft(1)

	metaStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("240"))

	starStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("220")) // gold

	urlStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("39")). // cyan-blue
			Italic(true)

	badgePublic = lipgloss.NewStyle().
			Foreground(lipgloss.Color("76")).
			Border(lipgloss.RoundedBorder()).
			PaddingLeft(1).PaddingRight(1)

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

		rank := fmt.Sprintf("#%d", i+1)
		header := lipgloss.JoinHorizontal(lipgloss.Center,
			headerStyle.Render(rank),
			repoNameStyle.Render(" "+item.FullName),
			"  ",
			badgePublic.Render(item.Visibility),
		)

		desc := descStyle.Render(item.Description)

		meta := metaStyle.Render(
			starStyle.Render("★") + fmt.Sprintf(" %-8d", item.StargazersCount) +
				"  🗓 " + formattedDate.Format("January 02, 2006") +
				"  " + urlStyle.Render(item.HTMLURL),
		)

		card := cardStyle.Width(width).Render(
			lipgloss.JoinVertical(lipgloss.Left, header, desc, meta),
		)

		fmt.Println(card)
	}

	fmt.Println(dividerStyle.Render(
		fmt.Sprintf("── %d results %s", len(items), strings.Repeat("─", width-14)),
	))
}
