package tui

import "github.com/charmbracelet/lipgloss"

var (
	listStyle              = lipgloss.NewStyle().Margin(1)
	statusMessageInfoStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#04B575", Dark: "#04B575"}).
				Render
	statusMessageErrorStyle = lipgloss.NewStyle().
				Foreground(lipgloss.AdaptiveColor{Light: "#FF0000", Dark: "#FF0000"}).
				Render
)

const (
	iconSelected    = "●"
	iconNotSelected = "○"
)
