package cmds

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/ssh"
	"github.com/gleich/terminal/internal/lcp"
	"github.com/gleich/terminal/internal/output"
)

func music(s ssh.Session, styles output.Styles) {
	var (
		headers = []string{"", "NAME", "ARTIST", "ALBUM"}
		data    [][]string
	)
	cacheData, err := lcp.FetchAppleMusicCache()
	if err != nil {
		fmt.Fprintln(s, styles.Red.Render("failed to load data from apple music cache"))
		return
	}
	rowStyle := lipgloss.NewStyle().Width(40)
	for i, s := range cacheData.Data.RecentlyPlayed {
		data = append(
			data,
			[]string{
				fmt.Sprint(i + 1),
				rowStyle.Render(s.Track),
				rowStyle.Render(s.Artist),
				rowStyle.Render(s.Album),
			},
		)
	}

	table := output.Table(styles).Headers(headers...).Rows(data...).Render()
	fmt.Fprintln(s)
	fmt.Fprintln(
		s,
		styles.Renderer.NewStyle().
			Width(lipgloss.Width(table)+10).
			Render("One of my favorite things in this world is music. Here are a few of the playlists I've built up over the last few years and my recently played songs. I am into everything from electronic to bossa nova. A few of my favorite artists are The Smiths, Coldplay, Daft Punk, and Earth Wind & Fire."),
	)

	fmt.Fprintln(s, "\nHere are 10 of my most recently played songs from Apple Music:")
	fmt.Fprintln(s, table)
	output.LiveFrom(s, styles, table, cacheData.Updated)
	fmt.Fprintln(s)
}
