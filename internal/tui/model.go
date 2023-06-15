package tui

import (
	"github.com/knipferrc/branch-cleaner/internal/config"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

// item represents a list item.
type item struct {
	title, desc string
	selected    bool
}

// Title returns the title of the list item.
func (i item) Title() string {
	if i.selected {
		return iconSelected + " " + i.title
	}

	return iconNotSelected + " " + i.title
}

// Description returns the description of the list item.
func (i item) Description() string { return i.desc }

// FilterValue returns the current filter value.
func (i item) FilterValue() string { return i.title }

// model represents the state of the UI.
type model struct {
	list      list.Model
	appConfig config.Config
}

// New creates an instance of the UI.
func New(cfg config.Config) model {
	l := list.New([]list.Item{}, list.NewDefaultDelegate(), 0, 0)
	l.Title = "Branch Cleaner"
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			deleteKey,
			selectKey,
		}
	}
	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			deleteKey,
			selectKey,
		}
	}

	return model{
		list:      l,
		appConfig: cfg,
	}
}
