package tui

import (
	"fmt"

	"github.com/caarlos0/env"
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

// config represents application wide configuration.
type config struct {
	ProtectedBranches []string `env:"PROTECTED_BRANCHES" envSeparator:":"`
}

// model represents the state of the UI.
type model struct {
	list      list.Model
	appConfig config
}

// New creates an instance of the UI.
func New() model {
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

	var cfg = config{
		ProtectedBranches: []string{"main", "master", "develop", "dev", "prod"},
	}

	if err := env.Parse(&cfg); err != nil {
		fmt.Printf("%+v\n", err)
	}

	return model{
		list:      l,
		appConfig: cfg,
	}
}
