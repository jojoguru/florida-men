package cliwriter

import (
	"fmt"

	"github.com/jojoguru/florida-men/domain"
	"github.com/pterm/pterm"
)

func Write(story *domain.Story) error {
	header := fmt.Sprintf("Florida Man story of the day [%s/%d]:", story.Date.Month, story.Date.Day)

	pterm.Println()
	pterm.Println(header)
	pterm.Println()

	titleStyle := pterm.NewStyle(pterm.FgGreen, pterm.Bold)
	titleStyle.Println("  " + story.Title)

	linkStyle := pterm.NewStyle(pterm.FgBlue)
	linkStyle.Print("  └─ Link: ")
	linkStyle.Add(*pterm.Italic.ToStyle(), *pterm.Bold.ToStyle()).Println(story.Link)
	pterm.Println()
	pterm.Println()

	return nil
}
