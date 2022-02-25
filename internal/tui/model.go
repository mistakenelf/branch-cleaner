package tui

import (
	"github.com/go-git/go-git/v5"
	"github.com/knipferrc/branch-cleaner/internal/config"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Bubble represents the state of the UI.
type Bubble struct {
	keys      keyMap
	help      help.Model
	spinner   spinner.Model
	viewport  viewport.Model
	list      list.Model
	appConfig config.Config
	repo      *git.Repository
	ready     bool
}

// NewBubble creates an instance of the UI.
func NewBubble() Bubble {
	cfg := config.GetConfig()
	keys := getDefaultKeyMap()

	s := spinner.New()
	s.Spinner = spinner.Dot

	h := help.New()
	h.Styles.FullKey.Foreground(lipgloss.Color("#ffffff"))
	h.Styles.FullDesc.Foreground(lipgloss.Color("#ffffff"))

	l := list.New(make([]list.Item, 0), list.NewDefaultDelegate(), 0, 0)
	l.Title = "Branches"

	return Bubble{
		keys:      keys,
		help:      h,
		spinner:   s,
		viewport:  viewport.Model{},
		list:      l,
		appConfig: cfg,
		ready:     false,
	}
}
