package tui

import (
	"github.com/charmbracelet/bubbles/key"
	tea "github.com/charmbracelet/bubbletea"
)

// Update handles updating the UI.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.screenWidth = msg.Width
		b.screenHeight = msg.Height
		v, h := listStyle.GetFrameSize()
		b.list.SetSize(msg.Width-h, msg.Height-v)
	case repoDataMsg:
		b.list.SetItems(msg)
		return b, nil
	case tea.KeyMsg:
		if key.Matches(msg, keyDeletedSelected) {
			var title string

			if i, ok := b.list.SelectedItem().(item); ok {
				title = i.Title()
			} else {
				return b, nil
			}

			b.list.NewStatusMessage(statusMessageStyle("Deleted " + title))
			cmds = append(cmds, deleteSelectedBranchCmd(title))
			index := b.list.Index()
			b.list.RemoveItem(index)

			if len(b.list.Items()) == 0 {
				keyDeletedSelected.SetEnabled(false)
			}
		}
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
