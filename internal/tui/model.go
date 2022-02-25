package tui

import (
	"github.com/knipferrc/branch-cleaner/internal/config"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Bubble represents the state of the UI.
type Bubble struct {
	spinner     spinner.Model
	list        list.Model
	appConfig   config.Config
	repo        *git.Repository
	previousKey tea.KeyMsg
	ready       bool
}

// NewBubble creates an instance of the UI.
func NewBubble() Bubble {
	cfg := config.GetConfig()

	s := spinner.New()
	s.Spinner = spinner.Dot

	l := list.New(make([]list.Item, 0), list.NewDefaultDelegate(), 0, 0)
	l.Title = "Branch Cleaner"

	return Bubble{
		spinner:   s,
		list:      l,
		appConfig: cfg,
		ready:     false,
	}
}
