package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/knipferrc/branch-cleaner/internal/config"
	"github.com/knipferrc/branch-cleaner/internal/tui"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "branch-cleaner",
	Short:   "branch-cleaner is a tui for cleaning up local git branches easily",
	Version: "0.0.2",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		config.LoadConfig()

		cfg := config.GetConfig()

		// If logging is enabled, logs will be output to debug.log.
		if cfg.Settings.EnableLogging {
			f, err := tea.LogToFile("debug.log", "debug")
			if err != nil {
				log.Fatal(err)
				os.Exit(1)
			}

			defer func() {
				if err = f.Close(); err != nil {
					log.Fatal(err)
					os.Exit(1)
				}
			}()
		}

		b := tui.NewBubble()
		var opts []tea.ProgramOption

		// Always append alt screen program option.
		opts = append(opts, tea.WithAltScreen())

		// Initialize new app.
		p := tea.NewProgram(b, opts...)
		if err := p.Start(); err != nil {
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
