package tui

import "github.com/charmbracelet/lipgloss"

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	return lipgloss.NewStyle().
		Width(b.screenWidth).
		Height(b.screenHeight).
		Render(b.list.View())
}
