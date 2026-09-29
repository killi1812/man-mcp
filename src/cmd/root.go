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
	Short: "man-mcp CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(app.Setup)
	rootCmd.AddCommand(version.VersionCmd)
}
