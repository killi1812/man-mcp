// Package version provides basic version information for app
package version

import (
	"fmt"

	"github.com/killi1812/man-mcp/app"
	"github.com/spf13/cobra"
)

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "prints app version, build type, commit hash, and build time stamp",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Version: %s\n", app.Version)
		fmt.Printf("Build: %s\n", app.Build)
		fmt.Printf("Commit: %s\n", app.CommitHash)
		fmt.Printf("Build Time Stamp: %s\n", app.BuildTimestamp)
	},
}

