package tui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles updating the UI.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.list.SetSize(msg.Width, msg.Height)

		if !b.ready {
			b.ready = true
		}

		return b, nil
	case repoDataMsg:
		var items []list.Item

		b.repo = msg.repo

		for _, branch := range msg.branches {
			items = append(items, item{
				title: branch.Name().Short(),
				desc:  branch.Hash().String(),
			})
		}

		b.list.SetItems(items)
		return b, nil
	case tea.KeyMsg:
		if b.previousKey.String() == "d" && msg.String() == "d" {
			return b, tea.Sequentially(b.deleteSelectedBranchCmd(), b.readCurrentGitBranchesCmd())
		}

		b.previousKey = msg
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)

	b.spinner, cmd = b.spinner.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
