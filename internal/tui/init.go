package tui

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
)

// Init initializes the UI.
func (b Bubble) Init() tea.Cmd {
	var cmds []tea.Cmd

	cmds = append(cmds, spinner.Tick)
	cmds = append(cmds, readCurrentGitBranchesCmd())

	return tea.Batch(cmds...)
}
