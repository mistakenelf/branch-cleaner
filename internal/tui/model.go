package tui

import (
	"github.com/knipferrc/branch-cleaner/internal/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

// item represents a list item.
type item struct {
	title, desc string
}

// Title returns the title of the list item.
func (i item) Title() string { return i.title }

// Description returns the description of the list item.
func (i item) Description() string { return i.desc }

// FilterValue returns the current filter value.
func (i item) FilterValue() string { return i.title }

// Bubble represents the state of the UI.
type Bubble struct {
	list         list.Model
	appConfig    config.Config
	screenWidth  int
	screenHeight int
}

// NewBubble creates an instance of the UI.
func NewBubble() Bubble {
	cfg := config.GetConfig()

	l := list.New(make([]list.Item, 0), list.NewDefaultDelegate(), 0, 0)
	l.Title = "Branch Cleaner"
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keyDeletedSelected,
		}
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			keyDeletedSelected,
		}
	}

	return Bubble{
		list:      l,
		appConfig: cfg,
	}
}
