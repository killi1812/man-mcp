---
description: OpenWolf protocol enforcement
---

# OpenWolf Protocol Rules

- To locate a symbol or file, run `openwolf find <name>` first (ranked shortlist, under 1k tokens). For one file's description and symbol ranges: `openwolf find --file <path>`. Never read `.wolf/anatomy.md` whole; it is an index.
- Check `.wolf/cerebrum.md` Do-Not-Repeat list before generating code (grep `"## Do-Not-Repeat"`); after a user correction, update `cerebrum.md` immediately.
- Do NOT manually update `.wolf/anatomy.md` or `.wolf/memory.md` unless your agent has no OpenWolf hooks installed.
- BEFORE fixing any bug: run `openwolf bug search "<error>"` or grep `.wolf/buglog.json`. AFTER fixing one: log it there (error_message, root_cause, fix, tags).
- When resuming a session, read `.wolf/STATUS.md` first; regenerate it with the `handoff` skill when a quest finishes.
