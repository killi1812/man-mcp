package mcp

import (
	"context"
	"testing"

	"github.com/killi1812/man-mcp/pkg/man"
	"github.com/mark3labs/mcp-go/mcp"
)

type mockClient struct{}

func (m *mockClient) Search(ctx context.Context, query string, section int) ([]man.SearchResult, error) {
	return []man.SearchResult{
		{Name: "ls", Section: "1", Description: "list directory contents"},
	}, nil
}

func (m *mockClient) Fetch(ctx context.Context, command string, section int) (string, error) {
	return `LS(1)                            User Commands                           LS(1)

NAME
     ls - list directory contents

OPTIONS
     -a, --all
            do not ignore entries starting with .
`, nil
}

func TestServerToolHandlers(t *testing.T) {
	svc := man.NewService(&mockClient{})
	srv := NewServer(svc)
	ctx := context.Background()

	testSearchHandler(t, srv, ctx)
	testPageHandler(t, srv, ctx)
	testSectionHandler(t, srv, ctx)
	testFlagHandler(t, srv, ctx)
	testListSectionsHandler(t, srv, ctx)
}

func testSearchHandler(t *testing.T, srv *Server, ctx context.Context) {
	req := makeCallRequest("man_search", map[string]any{"query": "ls"})
	res, err := srv.handleSearch(ctx, req)
	assertNilErr(t, err)
	assertFalse(t, res.IsError)
}

func testPageHandler(t *testing.T, srv *Server, ctx context.Context) {
	req := makeCallRequest("man_page", map[string]any{"command": "ls"})
	res, err := srv.handlePage(ctx, req)
	assertNilErr(t, err)
	assertFalse(t, res.IsError)
}

func testSectionHandler(t *testing.T, srv *Server, ctx context.Context) {
	req := makeCallRequest("man_section", map[string]any{
		"command":      "ls",
		"section_name": "NAME",
	})
	res, err := srv.handleSection(ctx, req)
	assertNilErr(t, err)
	assertFalse(t, res.IsError)
}

func testFlagHandler(t *testing.T, srv *Server, ctx context.Context) {
	req := makeCallRequest("man_flag", map[string]any{
		"command": "ls",
		"flag":    "-a",
	})
	res, err := srv.handleFlag(ctx, req)
	assertNilErr(t, err)
	assertFalse(t, res.IsError)
}

func testListSectionsHandler(t *testing.T, srv *Server, ctx context.Context) {
	req := makeCallRequest("man_list_sections", map[string]any{"command": "ls"})
	res, err := srv.handleListSections(ctx, req)
	assertNilErr(t, err)
	assertFalse(t, res.IsError)
}

func makeCallRequest(name string, args map[string]any) mcp.CallToolRequest {
	var req mcp.CallToolRequest
	req.Params.Name = name
	req.Params.Arguments = args
	return req
}

func assertNilErr(t *testing.T, err error) {
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
}

func assertFalse(t *testing.T, cond bool) {
	if cond {
		t.Fatal("expected false, got true")
	}
}
