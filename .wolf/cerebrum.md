---
description: learned preferences, project conventions, and Do-Not-Repeat rules
budget_tokens: 2000
---
# Cerebrum

> OpenWolf's learning memory. Updated automatically as the AI learns from interactions.
> Do not edit manually unless correcting an error.
> Last updated: 2026-09-29

## User Preferences

- Always use `task test` to run tests (never run `go -C src test` directly).
- Always use `task build` / `task dev` to build the project.

## Key Learnings

- **Project:** man-mcp
- **Description:** General terminal template to golang

## Do-Not-Repeat

- [2026-09-29] Do not run raw `go -C src test ...`. Always use `task test` (or `task coverage`) to execute tests.


## Decision Log

<!-- Significant technical decisions with rationale. Why X was chosen over Y. -->
