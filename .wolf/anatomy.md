# anatomy.md

> Auto-maintained by OpenWolf. Last scanned: 2026-09-29T10:53:38.341Z
> Files: 13 tracked | Anatomy hits: 0 | Misses: 0

> Project structure index. Auto-maintained by OpenWolf hooks and daemon.
> Run `openwolf scan` to generate, or wait for the first Claude Code session.
> Status: Pending initial scan

## ./

- `.gitignore` — Git ignore rules (~168 tok)
- `CLAUDE.md` — OpenWolf (~99 tok)
- `GEMINI.md` — OpenWolf (~75 tok)
- `LICENSE` — Project license (~9374 tok)
- `README.md` — Project documentation (~16 tok)
- `taskfile.yaml` (~463 tok)

## src/

- `go.mod` — Go module definition (~53 tok)
- `go.sum` — Go dependency checksums (~131 tok)
- `main.go` (~122 tok)

## src/app/

- `setup.go` — preforms basic app functions like setup, loading config and global definitions (~155 tok)
- `var.go` (~101 tok)
- `zap.go` — Declares name (~673 tok)

## src/cmd/version/

- `version.go` — provides basic version infromation for app (~293 tok)
