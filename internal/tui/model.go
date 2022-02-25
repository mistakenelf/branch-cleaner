package tui

import (
	"github.com/knipferrc/bubbletea-starter/internal/config"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

// Bubble represents the state of the UI.
type Bubble struct {
	keys      keyMap
	help      help.Model
	loader    spinner.Model
	viewport  viewport.Model
	appConfig config.Config
	ready     bool
}

// NewBubble creates an instance of the UI.
func NewBubble() Bubble {
	cfg := config.GetConfig()
	keys := getDefaultKeyMap()

	l := spinner.New()
	l.Spinner = spinner.Dot

	h := help.New()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	return Bubble{
		keys:      keys,
		help:      h,
		loader:    l,
		viewport:  viewport.Model{},
		appConfig: cfg,
		ready:     false,
	}
}
