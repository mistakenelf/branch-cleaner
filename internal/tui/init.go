package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the UI.
func (m model) Init() tea.Cmd {
	return tea.Batch(m.list.StartSpinner(), readCurrentGitBranchesCmd())
}
