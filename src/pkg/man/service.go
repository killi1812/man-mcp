package man

import (
	"context"
	"fmt"
	"sort"
	"sync"
)

// Service provides high-level man page operations with caching.
type Service struct {
	client Client
	cache  map[string]*Page
	mu     sync.RWMutex
}

// NewService creates a new Service.
func NewService(client Client) *Service {
	return &Service{
		client: client,
		cache:  make(map[string]*Page),
	}
}

// Search discovers commands matching a query via apropos.
func (s *Service) Search(ctx context.Context, query string, section int) ([]SearchResult, error) {
	return s.client.Search(ctx, query, section)
}

// GetPage retrieves and caches a parsed man page.
func (s *Service) GetPage(ctx context.Context, command string, section int) (*Page, error) {
	key := makeCacheKey(command, section)
	if page := s.readCache(key); page != nil {
		return page, nil
	}
	return s.fetchAndCache(ctx, command, section, key)
}

func makeCacheKey(cmd string, sec int) string {
	return fmt.Sprintf("%s:%d", cmd, sec)
}

func (s *Service) readCache(key string) *Page {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.cache[key]
}

func (s *Service) fetchAndCache(ctx context.Context, cmd string, sec int, key string) (*Page, error) {
	content, err := s.client.Fetch(ctx, cmd, sec)
	if err != nil {
		return nil, err
	}
	page := buildPage(cmd, sec, content)
	s.writeCache(key, page)
	return page, nil
}

func buildPage(cmd string, sec int, content string) *Page {
	return &Page{
		Command:    cmd,
		Section:    sec,
		RawContent: content,
		Sections:   ParseSections(content),
	}
}

func (s *Service) writeCache(key string, page *Page) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.cache[key] = page
}

// GetSection extracts a specific section from a man page.
func (s *Service) GetSection(ctx context.Context, cmd string, secName string, sec int) (string, error) {
	page, err := s.GetPage(ctx, cmd, sec)
	if err != nil {
		return "", err
	}
	return extractSectionFromPage(page, secName)
}

func extractSectionFromPage(p *Page, name string) (string, error) {
	val, ok := ExtractSection(p.RawContent, name)
	if !ok {
		return "", fmt.Errorf("section %q not found in %s", name, p.Command)
	}
	return val, nil
}

// GetOption extracts a flag definition from a man page.
func (s *Service) GetOption(ctx context.Context, cmd string, flag string, sec int) (string, error) {
	page, err := s.GetPage(ctx, cmd, sec)
	if err != nil {
		return "", err
	}
	return extractOptionFromPage(page, flag)
}

func extractOptionFromPage(p *Page, flag string) (string, error) {
	val, ok := ExtractFlag(p.RawContent, flag)
	if !ok {
		return "", fmt.Errorf("flag %q not found in %s", flag, p.Command)
	}
	return val, nil
}

// ListSections returns all section headers available in a page.
func (s *Service) ListSections(ctx context.Context, cmd string, sec int) ([]string, error) {
	page, err := s.GetPage(ctx, cmd, sec)
	if err != nil {
		return nil, err
	}
	return sortSectionKeys(page.Sections), nil
}

func sortSectionKeys(m map[string]string) []string {
	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}
