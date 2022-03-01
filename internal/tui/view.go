package tui

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	return listStyle.Render(b.list.View())
}
