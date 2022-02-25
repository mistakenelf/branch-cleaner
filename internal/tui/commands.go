package tui

import (
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

type errorMsg error
type repoDataMsg struct {
	repo     *git.Repository
	branches []*plumbing.Reference
}

func (b Bubble) readCurrentGitBranchesCmd() tea.Cmd {
	return func() tea.Msg {
		r, err := git.PlainOpen(".git")
		if err != nil {
			return errorMsg(err)
		}

		branches, err := r.Branches()
		if err != nil {
			return errorMsg(err)
		}

		var branchNames []*plumbing.Reference
		err = branches.ForEach(func(b *plumbing.Reference) error {
			branchNames = append(branchNames, b)

			return nil
		})

		if err != nil {
			return errorMsg(err)
		}

		return repoDataMsg{
			repo:     r,
			branches: branchNames,
		}
	}
}

func (b Bubble) deleteSelectedBranchCmd() tea.Cmd {
	return func() tea.Msg {
		var selectedBranch string

		if i, ok := b.list.SelectedItem().(item); ok {
			selectedBranch = i.Title()
		} else {
			return nil
		}

		err := b.repo.DeleteBranch(selectedBranch)
		if err != nil {
			return errorMsg(err)
		}

		return nil
	}
}
