package tui

import (
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

		b.list.SetSize(msg.Width-appStyle.GetHorizontalPadding(), msg.Height-appStyle.GetVerticalPadding())
	case repoDataMsg:
		b.list.SetItems(msg)
		return b, nil
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
