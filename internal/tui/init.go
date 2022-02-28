package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the UI.
func (b Bubble) Init() tea.Cmd {
	return tea.Batch(b.list.StartSpinner(), readCurrentGitBranchesCmd())
}
