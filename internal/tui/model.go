package tui

import (
	"github.com/knipferrc/branch-cleaner/internal/config"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/spinner"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Bubble represents the state of the UI.
type Bubble struct {
	spinner      spinner.Model
	list         list.Model
	delegateKeys *delegateKeyMap
	appConfig    config.Config
	ready        bool
}

// NewBubble creates an instance of the UI.
func NewBubble() Bubble {
	cfg := config.GetConfig()

	s := spinner.New()
	s.Spinner = spinner.Dot

	delegateKeys := newDelegateKeyMap()
	delegate := newItemDelegate(delegateKeys)
	l := list.New(make([]list.Item, 0), delegate, 0, 0)
	l.Title = "Branch Cleaner"

	return Bubble{
		spinner:      s,
		list:         l,
		appConfig:    cfg,
		delegateKeys: delegateKeys,
		ready:        false,
	}
}
