package man

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
)

var aproposLineRegex = regexp.MustCompile(`^(\S+)\s+\(([^\)]+)\)\s+-\s+(.*)$`)

// SystemClient runs OS man and apropos commands.
type SystemClient struct{}

// NewSystemClient creates a new SystemClient.
func NewSystemClient() *SystemClient {
	return &SystemClient{}
}

// Search queries apropos for commands matching query.
func (c *SystemClient) Search(ctx context.Context, query string, section int) ([]SearchResult, error) {
	args := buildAproposArgs(query, section)
	cmd := exec.CommandContext(ctx, "apropos", args...)
	out, err := runCommand(cmd)
	if err != nil {
		return handleAproposError(err)
	}
	return parseAproposOutput(out), nil
}

func buildAproposArgs(query string, section int) []string {
	if section > 0 {
		return []string{"-s", strconv.Itoa(section), query}
	}
	return []string{query}
}

func handleAproposError(err error) ([]SearchResult, error) {
	if isExitCode1(err) {
		return []SearchResult{}, nil
	}
	return nil, err
}

func isExitCode1(err error) bool {
	if exitErr, ok := err.(*exec.ExitError); ok {
		return exitErr.ExitCode() == 1
	}
	return false
}

func parseAproposOutput(out string) []SearchResult {
	lines := strings.Split(out, "\n")
	var results []SearchResult
	for _, line := range lines {
		results = appendMatch(results, line)
	}
	return results
}

func appendMatch(results []SearchResult, line string) []SearchResult {
	m := aproposLineRegex.FindStringSubmatch(strings.TrimSpace(line))
	if len(m) != 4 {
		return results
	}
	return append(results, SearchResult{
		Name:        m[1],
		Section:     m[2],
		Description: m[3],
	})
}

// Fetch executes man to retrieve formatted man page text.
func (c *SystemClient) Fetch(ctx context.Context, command string, section int) (string, error) {
	args := buildManArgs(command, section)
	cmd := exec.CommandContext(ctx, "man", args...)
	setManEnv(cmd)
	out, err := runCommand(cmd)
	if err != nil {
		return "", fmt.Errorf("man page not found for %s: %w", command, err)
	}
	return CleanFormatting(out), nil
}

func buildManArgs(command string, section int) []string {
	if section > 0 {
		return []string{"-P", "cat", strconv.Itoa(section), command}
	}
	return []string{"-P", "cat", command}
}

func setManEnv(cmd *exec.Cmd) {
	cmd.Env = append(cmd.Environ(), "MANPAGER=cat", "PAGER=cat")
}

func runCommand(cmd *exec.Cmd) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return stdout.String(), nil
}
