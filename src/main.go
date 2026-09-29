package main

import (
	"github.com/killi1812/man-mcp/app"
	"github.com/killi1812/man-mcp/cmd"
)

func init() {
	app.Setup()
}

func main() {
	cmd.Execute()
}
