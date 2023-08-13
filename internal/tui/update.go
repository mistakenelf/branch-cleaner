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
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		v, h := listStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v)
	case repoDataMsg:
		m.list.SetItems(msg)

		return m, nil
	case errorMsg:
		return m, m.list.NewStatusMessage(statusMessageErrorStyle(msg.Error()))
	case tea.KeyMsg:
		if m.list.FilterState() == list.Filtering {
			break
		}

		switch {
		case key.Matches(msg, deleteKey):
			var deletedBranchTitles []string
			var protectedBranchTitles []string

			selected, unselected := splitBySelection(m.list.Items())

			for _, it := range selected {
				if !contains(m.appConfig.ProtectedBranches, it.title) {
					cmds = append(cmds, deleteSelectedBranchCmd(it.title))
					deletedBranchTitles = append(deletedBranchTitles, it.title)

					if len(m.list.Items()) == 0 {
						deleteKey.SetEnabled(false)
					}
				} else {
					protectedBranchTitles = append(protectedBranchTitles, it.title)
				}
			}

			if len(deletedBranchTitles) > 0 {
				statusMessage := fmt.Sprintf("Deleted branches: %s", strings.Join(deletedBranchTitles, ", "))
				statusCmd := m.list.NewStatusMessage(
					statusMessageInfoStyle(statusMessage),
				)
				cmds = append(cmds, statusCmd)
				m.list.ResetSelected()
			}

			if len(protectedBranchTitles) > 0 {
				statusMessage := fmt.Sprintf("Cannot delete protected branch(es): %s", strings.Join(protectedBranchTitles, ", "))
				statusCmd := m.list.NewStatusMessage(
					statusMessageErrorStyle(statusMessage),
				)
				return m, statusCmd
			}

			if len(unselected) > 0 {
				var items = make([]list.Item, 0, len(unselected))
				for _, unselectedItem := range unselected {
					items = append(items, item{
						title: unselectedItem.title,
						desc:  unselectedItem.desc,
					})
				}

				m.list.SetItems(items)
			}
		case key.Matches(msg, selectAllKey):
			for idx, i := range m.list.Items() {
				item := i.(item)
				if !contains(m.appConfig.ProtectedBranches, item.title) {
					item.selected = true
					m.list.RemoveItem(idx)
					cmds = append(cmds, m.list.InsertItem(idx, item))
				}
			}
		case key.Matches(msg, unselectAllKey):
			for idx, i := range m.list.Items() {
				item := i.(item)
				item.selected = false
				m.list.RemoveItem(idx)
				cmds = append(cmds, m.list.InsertItem(idx, item))
			}
		case key.Matches(msg, selectKey):
			idx := m.list.Index()
			item := m.list.SelectedItem().(item)
			item.selected = !item.selected
			m.list.RemoveItem(idx)
			cmds = append(cmds, m.list.InsertItem(idx, item))
		}
	}

	m.list, cmd = m.list.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
