package cmd

import (
	"fmt"
	"os"

	"github.com/killi1812/man-mcp/app"
	"github.com/killi1812/man-mcp/cmd/version"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "man-mcp",
	Short: "man-mcp is a Model Context Protocol server exposing system man pages",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServe()
	},
}

// Execute runs the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(app.Setup)
	addPersistentFlags()
	registerSubcommands()
}

func addPersistentFlags() {
	rootCmd.PersistentFlags().StringVar(&app.LogFilePath, "log", "", "Path to log file")
	rootCmd.PersistentFlags().BoolVarP(&app.Verbose, "verbose", "v", false, "Enable verbose/debug logging")
}

func registerSubcommands() {
	rootCmd.AddCommand(serveCmd)
	rootCmd.AddCommand(version.VersionCmd)
}
