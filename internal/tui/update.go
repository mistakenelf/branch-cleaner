package tui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles updating the UI.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.viewport.Height = msg.Height
		b.viewport.Width = msg.Width
		b.help.Width = msg.Width
		b.list.SetSize(msg.Width, msg.Height)

		if !b.ready {
			b.ready = true
		}

		return b, nil
	case errorMsg:
		b.viewport.SetContent(msg.Error())
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

		switch {
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		case key.Matches(msg, b.keys.Help):
			b.help.ShowAll = !b.help.ShowAll

			return b, nil
		}

		b.previousKey = msg
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)

	b.spinner, cmd = b.spinner.Update(msg)
	cmds = append(cmds, cmd)

	b.viewport, cmd = b.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
