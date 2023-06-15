package tui

// View returns a string representation of the entire application UI.
func (m model) View() string {
	return listStyle.Render(m.list.View())
}
