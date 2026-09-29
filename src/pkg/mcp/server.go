package mcp

import (
	"context"
	"encoding/json"

	"github.com/killi1812/man-mcp/app"
	"github.com/killi1812/man-mcp/pkg/man"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Server encapsulates the MCP server for man-mcp.
type Server struct {
	mcpServer *server.MCPServer
	service   *man.Service
}

// NewServer initializes the man-mcp server and registers tools.
func NewServer(service *man.Service) *Server {
	s := &Server{
		mcpServer: server.NewMCPServer("man-mcp", app.Version),
		service:   service,
	}
	s.registerAllTools()
	return s
}

func (s *Server) registerAllTools() {
	s.registerSearchTool()
	s.registerPageTool()
	s.registerSectionTool()
	s.registerFlagTool()
	s.registerListSectionsTool()
}

func (s *Server) registerSearchTool() {
	tool := mcp.NewTool("man_search",
		mcp.WithDescription("Search man pages using apropos keywords to discover relevant commands and utilities"),
		mcp.WithString("query", mcp.Required(), mcp.Description("Keyword or search term")),
		mcp.WithInteger("section", mcp.Description("Optional man section filter (e.g. 1 for user commands, 2 for syscalls, 3 for libc)")),
	)
	s.mcpServer.AddTool(tool, s.handleSearch)
}

func (s *Server) handleSearch(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	query, err := req.RequireString("query")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sec := req.GetInt("section", 0)
	results, err := s.service.Search(ctx, query, sec)
	return formatSearchResults(results, err)
}

func formatSearchResults(res []man.SearchResult, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	data, _ := json.MarshalIndent(res, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func (s *Server) registerPageTool() {
	tool := mcp.NewTool("man_page",
		mcp.WithDescription("Retrieve the full rendered text of a manual page"),
		mcp.WithString("command", mcp.Required(), mcp.Description("Name of command, library call, or topic")),
		mcp.WithInteger("section", mcp.Description("Optional manual section number (e.g. 1, 2, 3)")),
	)
	s.mcpServer.AddTool(tool, s.handlePage)
}

func (s *Server) handlePage(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cmd, err := req.RequireString("command")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sec := req.GetInt("section", 0)
	page, err := s.service.GetPage(ctx, cmd, sec)
	return formatPageResult(page, err)
}

func formatPageResult(p *man.Page, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(p.RawContent), nil
}

func (s *Server) registerSectionTool() {
	tool := mcp.NewTool("man_section",
		mcp.WithDescription("Extract a specific section from a manual page (e.g. OPTIONS, EXAMPLES, DESCRIPTION, SYNOPSIS) to save context tokens"),
		mcp.WithString("command", mcp.Required(), mcp.Description("Name of command or topic")),
		mcp.WithString("section_name", mcp.Required(), mcp.Description("Section heading name (e.g. OPTIONS, EXAMPLES, DESCRIPTION, SYNOPSIS, EXIT STATUS)")),
		mcp.WithInteger("section", mcp.Description("Optional manual section number")),
	)
	s.mcpServer.AddTool(tool, s.handleSection)
}

func (s *Server) handleSection(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cmd, secName, err := extractCmdAndSecName(req)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sec := req.GetInt("section", 0)
	val, err := s.service.GetSection(ctx, cmd, secName, sec)
	return formatTextResult(val, err)
}

func extractCmdAndSecName(req mcp.CallToolRequest) (string, string, error) {
	cmd, err := req.RequireString("command")
	if err != nil {
		return "", "", err
	}
	secName, err := req.RequireString("section_name")
	if err != nil {
		return "", "", err
	}
	return cmd, secName, nil
}

func (s *Server) registerFlagTool() {
	tool := mcp.NewTool("man_flag",
		mcp.WithDescription("Extract the definition and documentation for a specific CLI flag or option (e.g. -a, --all, -l)"),
		mcp.WithString("command", mcp.Required(), mcp.Description("Name of command")),
		mcp.WithString("flag", mcp.Required(), mcp.Description("Flag or option name (e.g. -a, --all, -l)")),
		mcp.WithInteger("section", mcp.Description("Optional manual section number")),
	)
	s.mcpServer.AddTool(tool, s.handleFlag)
}

func (s *Server) handleFlag(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cmd, flag, err := extractCmdAndFlag(req)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sec := req.GetInt("section", 0)
	val, err := s.service.GetOption(ctx, cmd, flag, sec)
	return formatTextResult(val, err)
}

func extractCmdAndFlag(req mcp.CallToolRequest) (string, string, error) {
	cmd, err := req.RequireString("command")
	if err != nil {
		return "", "", err
	}
	flag, err := req.RequireString("flag")
	if err != nil {
		return "", "", err
	}
	return cmd, flag, nil
}

func (s *Server) registerListSectionsTool() {
	tool := mcp.NewTool("man_list_sections",
		mcp.WithDescription("List available section headings for a command's man page"),
		mcp.WithString("command", mcp.Required(), mcp.Description("Name of command")),
		mcp.WithInteger("section", mcp.Description("Optional manual section number")),
	)
	s.mcpServer.AddTool(tool, s.handleListSections)
}

func (s *Server) handleListSections(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	cmd, err := req.RequireString("command")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	sec := req.GetInt("section", 0)
	sections, err := s.service.ListSections(ctx, cmd, sec)
	return formatStringSliceResult(sections, err)
}

func formatStringSliceResult(sections []string, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	data, _ := json.MarshalIndent(sections, "", "  ")
	return mcp.NewToolResultText(string(data)), nil
}

func formatTextResult(val string, err error) (*mcp.CallToolResult, error) {
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}
	return mcp.NewToolResultText(val), nil
}

// ServeStdio launches the server listening on stdio.
func (s *Server) ServeStdio() error {
	return server.ServeStdio(s.mcpServer)
}
