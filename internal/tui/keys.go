package tui

import "github.com/charmbracelet/bubbles/key"

var (
	keyDeletedSelected = key.NewBinding(key.WithKeys("x"), key.WithHelp("x", "delete selected branch"))
)
