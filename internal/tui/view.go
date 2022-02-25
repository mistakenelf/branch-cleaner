package tui

import (
	"fmt"
)

// View returns a string representation of the entire application UI.
func (b Bubble) View() string {
	if !b.ready {
		return fmt.Sprintf("%s%s", b.spinner.View(), "loading...")
	}

	return b.list.View()
}
