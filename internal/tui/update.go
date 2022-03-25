package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

// contains returns true if the slice contains the string.
func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// splitBySelections splits the items into two slices, one with selected items and one with unselected items.
func splitBySelection(items []list.Item) ([]item, []item) {
	var selected, unselected []item
	for _, it := range items {
		item := it.(item)
		if item.selected {
			selected = append(selected, item)
		} else {
			unselected = append(unselected, item)
		}
	}
	return selected, unselected
}

// Update handles updating the UI.
func (b Bubble) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		v, h := listStyle.GetFrameSize()
		b.list.SetSize(msg.Width-h, msg.Height-v)
	case repoDataMsg:
		b.list.SetItems(msg)

		return b, nil
	case errorMsg:
		return b, b.list.NewStatusMessage(statusMessageErrorStyle(msg.Error()))
	case tea.KeyMsg:
		if b.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, deleteKey):
			var deletedBranchTitles []string
			var protectedBranchTitles []string

			selected, unselected := splitBySelection(b.list.Items())

			for _, it := range selected {
				if !contains(b.appConfig.Settings.ProtectedBranches, it.title) {
					cmds = append(cmds, deleteSelectedBranchCmd(it.title))
					deletedBranchTitles = append(deletedBranchTitles, it.title)

					if len(b.list.Items()) == 0 {
						deleteKey.SetEnabled(false)
					}
				} else {
					protectedBranchTitles = append(protectedBranchTitles, it.title)
				}
			}

			if len(deletedBranchTitles) > 0 {
				statusMessage := fmt.Sprintf("Deleted branches: %s", strings.Join(deletedBranchTitles, ", "))
				statusCmd := b.list.NewStatusMessage(
					statusMessageInfoStyle(statusMessage),
				)
				cmds = append(cmds, statusCmd)
				b.list.ResetSelected()
			}

			if len(protectedBranchTitles) > 0 {
				statusMessage := fmt.Sprintf("Cannot delete protected branch(es): %s", strings.Join(protectedBranchTitles, ", "))
				statusCmd := b.list.NewStatusMessage(
					statusMessageErrorStyle(statusMessage),
				)
				return b, statusCmd
			}

			if len(unselected) > 0 {
				var items = make([]list.Item, 0, len(unselected))
				for _, unselectedItem := range unselected {
					items = append(items, item{
						title: unselectedItem.title,
						desc:  unselectedItem.desc,
					})
				}

				b.list.SetItems(items)
			}
		case key.Matches(msg, selectAllKey):
			for idx, i := range b.list.Items() {
				item := i.(item)
				if !contains(b.appConfig.Settings.ProtectedBranches, item.title) {
					item.selected = true
					b.list.RemoveItem(idx)
					cmds = append(cmds, b.list.InsertItem(idx, item))
				}
			}
		case key.Matches(msg, unselectAllKey):
			for idx, i := range b.list.Items() {
				item := i.(item)
				item.selected = false
				b.list.RemoveItem(idx)
				cmds = append(cmds, b.list.InsertItem(idx, item))
			}
		case key.Matches(msg, selectKey):
			idx := b.list.Index()
			item := b.list.SelectedItem().(item)
			item.selected = !item.selected
			b.list.RemoveItem(idx)
			cmds = append(cmds, b.list.InsertItem(idx, item))
		}
	}

	b.list, cmd = b.list.Update(msg)
	cmds = append(cmds, cmd)

	return b, tea.Batch(cmds...)
}
