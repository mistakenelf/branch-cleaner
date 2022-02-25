package tui

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
)

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	var currentView string

	if !b.ready {
		return fmt.Sprintf("%s%s", b.loader.View(), "loading...")
	}

	if b.help.ShowAll {
		currentView = b.help.View(b.keys)
	} else {
		currentView = b.viewport.View()
	}

	return lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFFFFF")).
		Bold(true).
		Italic(true).
		Render(currentView)
}
