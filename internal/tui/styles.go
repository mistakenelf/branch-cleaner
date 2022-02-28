package tui

import "github.com/charmbracelet/lipgloss"

var (
	listStyle          = lipgloss.NewStyle().Margin(2)
	statusMessageStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
)
