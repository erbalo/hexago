package menu

import "github.com/charmbracelet/lipgloss"

const (
	// Define colors
	darkGray    = lipgloss.Color("#767676")
	awesomeBlue = lipgloss.Color("#288BA8")
	pink        = lipgloss.Color("#E389B9")
)

var (
	// Define styles
	cursorStyle   = lipgloss.NewStyle().Foreground(awesomeBlue).Bold(true)
	choiceStyle   = lipgloss.NewStyle().Foreground(pink)
	continueStyle = lipgloss.NewStyle().Foreground(darkGray)
)
