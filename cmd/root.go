package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/mistakenelf/branch-cleaner/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "branch-cleaner",
	Short:   "branch-cleaner is a tui for cleaning up local git branches easily",
	Version: "1.0.1",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Getenv("DEBUG")) > 0 {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				fmt.Println("fatal:", err)
				os.Exit(1)
			}

			defer f.Close()
		}

		m := tui.New()
		var opts []tea.ProgramOption

		// Always append alt screen program option.
		opts = append(opts, tea.WithAltScreen())

		// Initialize new app.
		p := tea.NewProgram(m, opts...)
		if _, err := p.Run(); err != nil {
			log.Fatal("Failed to start branch-cleaner", err)
			os.Exit(1)
		}
	},
}

// Execute executes the root command which starts the application.
func Execute() {
	rootCmd.AddCommand(updateCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
