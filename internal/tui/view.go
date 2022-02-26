package tui

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	return b.list.View()
}
