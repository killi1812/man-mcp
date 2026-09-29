package cmd

import (
	"github.com/killi1812/man-mcp/pkg/man"
	"github.com/killi1812/man-mcp/pkg/mcp"
	"github.com/spf13/cobra"
)

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the man-mcp server over standard I/O",
	RunE: func(cmd *cobra.Command, args []string) error {
		return runServe()
	},
}

func runServe() error {
	client := man.NewSystemClient()
	svc := man.NewService(client)
	server := mcp.NewServer(svc)
	return server.ServeStdio()
}
