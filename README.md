# man-mcp

A [Model Context Protocol (MCP)](https://modelcontextprotocol.io/) server in Go that gives AI assistants, LLMs, and agents direct access to local system manual pages (`man`) as an authoritative, zero-hallucination source of truth.
[![M8ven Score](https://m8ven.ai/badge/mcp/killi1812-man-mcp-1b209g?v=dfab5a2fd742d32da4035d61c7a9a5e9)](https://m8ven.ai/mcp/killi1812-man-mcp-1b209g?s=readme)
---

## MCP Tools

| Tool | Parameters | Description |
| :--- | :--- | :--- |
| `man_search` | `query` (string, req), `section` (int, opt) | Search manual descriptions using `apropos` / `man -k`. |
| `man_page` | `command` (string, req), `section` (int, opt) | Retrieve the full rendered manual page. |
| `man_section` | `command` (string, req), `section_name` (string, req), `section` (int, opt) | Extract a specific section (e.g. `OPTIONS`, `EXAMPLES`, `SYNOPSIS`). |
| `man_flag` | `command` (string, req), `flag` (string, req), `section` (int, opt) | Extract the definition and documentation for a specific CLI flag (e.g. `-l`, `--all`). |
| `man_list_sections`| `command` (string, req), `section` (int, opt) | List available section headings for a command. |

---

## Installation & Build

### Prerequisites
- Go 1.24+
- `man-db` / `man` installed on system
- [Task](https://taskfile.dev/) (recommended)

### Using Task

```bash
# Build the binary
task build

# Build for development (includes dev ldflags)
task dev

# Run unit tests
task test

# Generate test coverage report
task coverage

# Install binary to $GOBIN or ~/.local/bin
task install

# Clean build artifacts
task clean
```

### Direct Go Build

```bash
cd src
go build -o ../man-mcp .
```

---

## Configuration

Add `man-mcp` to your MCP client configuration (e.g. Claude Desktop, Cursor, Gemini CLI, Antigravity, or VS Code):

### Claude Desktop (`claude_desktop_config.json`)

```json
{
  "mcpServers": {
    "man": {
      "command": "man-mcp",
      "args": ["serve"]
    }
  }
}
```

### With Custom Logging

```json
{
  "mcpServers": {
    "man": {
      "command": "man-mcp",
      "args": ["serve", "--log", "/tmp/man-mcp.log", "--verbose"]
    }
  }
}
```

---

## CLI Usage

```text
man-mcp is a Model Context Protocol server exposing system man pages

Usage:
  man-mcp [flags]
  man-mcp [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  help        Help about any command
  serve       Start the man-mcp server over standard I/O (default action)
  version     prints app version, build type, commit hash, and build time stamp

Flags:
  -h, --help         help for man-mcp
      --log string   Path to log file
  -v, --verbose      Enable verbose/debug logging
```

---

## Packaging

TODO: package

Distribution packaging templates are available in [`packaging/`](packaging/):
- **Arch Linux (AUR)**: [`packaging/aur/PKGBUILD`](packaging/aur/PKGBUILD)
- **Fedora / RHEL (RPM)**: [`packaging/fedora/man-mcp.spec`](packaging/fedora/man-mcp.spec)
- **macOS / Linuxbrew**: [`packaging/homebrew/man-mcp.rb`](packaging/homebrew/man-mcp.rb)

---

## License

This project is licensed under the terms of the GNU General Public License v3.0 or later ([LICENSE](LICENSE)).
