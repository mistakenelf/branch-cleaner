package tui

import (
	"github.com/knipferrc/branch-cleaner/internal/config"

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
	delegateKeys *delegateKeyMap
	appConfig    config.Config
}

// NewBubble creates an instance of the UI.
func NewBubble() Bubble {
	cfg := config.GetConfig()

	delegateKeys := newDelegateKeyMap()
	delegate := newItemDelegate(delegateKeys)
	l := list.New(make([]list.Item, 0), delegate, 0, 0)
	l.Title = "Branch Cleaner"

	return Bubble{
		list:         l,
		appConfig:    cfg,
		delegateKeys: delegateKeys,
	}
}
