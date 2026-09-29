package man

import (
	"context"
	"errors"
	"testing"
)

type mockClient struct {
	fetchCalls int
	content    string
	results    []SearchResult
}

func (m *mockClient) Search(ctx context.Context, query string, section int) ([]SearchResult, error) {
	return m.results, nil
}

func (m *mockClient) Fetch(ctx context.Context, command string, section int) (string, error) {
	m.fetchCalls++
	if m.content == "" {
		return "", errors.New("not found")
	}
	return m.content, nil
}

func TestServiceGetPageAndCache(t *testing.T) {
	mock := &mockClient{content: sampleManPage}
	svc := NewService(mock)
	ctx := context.Background()

	page1, err1 := svc.GetPage(ctx, "ls", 1)
	assertTrue(t, err1 == nil)
	assertMatch(t, page1.Command, "ls")

	page2, err2 := svc.GetPage(ctx, "ls", 1)
	assertTrue(t, err2 == nil)
	assertTrue(t, page1 == page2)
	assertTrue(t, mock.fetchCalls == 1)
}

func TestServiceGetSection(t *testing.T) {
	mock := &mockClient{content: sampleManPage}
	svc := NewService(mock)
	ctx := context.Background()

	desc, err := svc.GetSection(ctx, "ls", "DESCRIPTION", 1)
	assertTrue(t, err == nil)
	assertContainsString(t, desc, "List information about the FILEs.")

	_, errNotFound := svc.GetSection(ctx, "ls", "UNKNOWN", 1)
	assertTrue(t, errNotFound != nil)
}

func TestServiceGetOption(t *testing.T) {
	mock := &mockClient{content: sampleManPage}
	svc := NewService(mock)
	ctx := context.Background()

	opt, err := svc.GetOption(ctx, "ls", "-a", 1)
	assertTrue(t, err == nil)
	assertContainsString(t, opt, "do not ignore entries")
}

func TestServiceListSections(t *testing.T) {
	mock := &mockClient{content: sampleManPage}
	svc := NewService(mock)
	ctx := context.Background()

	sections, err := svc.ListSections(ctx, "ls", 1)
	assertTrue(t, err == nil)
	assertTrue(t, len(sections) >= 4)
}
