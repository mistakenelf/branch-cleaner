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
		b.viewport.Height = msg.Height
		b.viewport.Width = msg.Width
		b.help.Width = msg.Width
		b.viewport.SetContent("Welcome to the bubbletea-starter app")

		if !b.ready {
			b.ready = true
		}

	case tea.KeyMsg:
		switch {
		case key.Matches(msg, b.keys.Quit):
			return b, tea.Quit
		case key.Matches(msg, b.keys.Help):
			b.help.ShowAll = !b.help.ShowAll

			return b, nil
		}
	}

	b.loader, cmd = b.loader.Update(msg)
	cmds = append(cmds, cmd)

	b.viewport, cmd = b.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
