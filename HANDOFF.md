# Handoff: Blog Post — "A Weekend with OpenClaw"

**Generated**: 2026-02-08
**Branch**: main
**Status**: In Progress

## Goal

Help Niko write his first blog post for nkdem.net about setting up OpenClaw (AI assistant framework) on his Mac mini over a weekend. The site infrastructure (typography, tag pages, layouts) is done — the blog post content still needs writing.

## Completed

- [x] Extracted structured outline from raw notes (SSH'd to Mac mini, read 3 memory files)
- [x] Produced timeline, technical inventory, key moments, lessons learned, stats
- [x] Refined outline based on Niko's input on motivation/framing/tone
- [x] Installed `@tailwindcss/typography` — prose styling for markdown blog content
- [x] Added `prose`/`prose-invert` wrapper to `base.njk` with opt-out via `prose: false` frontmatter
- [x] Restyled blog listing (`/blog/`) — pill-shaped tag badges, proper spacing, inline dates
- [x] Auto-generated per-topic tag pages via Eleventy pagination (`src/blog/tags/topic.njk`)
- [x] Tags displayed on article pages (between title and content)
- [x] Fixed homepage prose styling (`prose: false`)
- [x] Agreed on blog post structure: Intro → The Setup → Things that Broke → Privacy & Security → What's Next
- [x] Intro drafted and reviewed
- [x] Description finalised

## Not Yet Done

- [ ] Write "The Setup" section — Nix, Tailscale, whisper-cpp, dummy HDMI plug, dead ends (sherpa-onnx, Ollama)
- [ ] Write "Things that Broke" section — model ID retry storm, 3am Strava DM, Signal Base64 bug, Peekaboo vs TCC
- [ ] Write "Privacy & Security" section — Tailscale helps network layer, LLM data exposure is the real concern, venice.ai, self-hosted LLM future
- [ ] Write "What's Next" section — ongoing project, Forgejo, Vaultwarden, self-hosted LLM, reignited self-hosting spark
- [ ] Remove placeholder code examples (Code Example, Terminal Example sections at bottom of first-post.md) before publishing
- [ ] Set `draft: false` when ready to publish

## Failed Approaches (Don't Repeat These)

- **`date: Last Modified` in frontmatter** — Eleventy resolves this to the file's mtime, but the custom `date` filter calls `.toISOString()` and `sortEntriesByYear` does `new Date(entry.date)`, which caused the blog listing to show "No entries yet". Fixed by using explicit date `2026-02-08`.
- **Prose wrapper on all pages** — Adding `prose` class globally in `base.njk` broke the homepage and blog listing (unwanted bullets, overridden spacing). Fixed with `prose: false` opt-out in frontmatter.

## Key Decisions

| Decision | Rationale |
|----------|-----------|
| `@tailwindcss/typography` over hand-written CSS | Less work, sensible defaults, dark mode support via `prose-invert` |
| Prose opt-out (`prose: false`) rather than opt-in | Blog posts (markdown) are the main prose content; fewer pages need opt-out than opt-in |
| Auto-generated tag pages via pagination | No manual work when adding new topics — just add to frontmatter |
| 4 sections after intro | Each has a clear purpose: what was built, what broke, the trust question, what's next |
| `draft: true` | Post is not ready for production builds yet |

## Current State

**Working**: Dev server at `localhost:8081`, blog listing with tag pills, per-topic tag pages, prose styling on articles, tags shown on article pages

**Broken**: Nothing

**Uncommitted Changes**: `src/blog/entry/first-post.md` — intro text, section headings, updated frontmatter (title, topics, description, date). Intentionally uncommitted (draft content).

## Files to Know

| File | Why It Matters |
|------|----------------|
| `src/blog/entry/first-post.md` | The blog post being written |
| `src/layouts/base.njk` | Base layout — prose wrapper, topic tags display |
| `src/blog/index.njk` | Blog listing page with tag pills |
| `src/blog/tags/topic.njk` | Auto-generated per-topic tag pages (Eleventy pagination) |
| `eleventy.config.js` | `topics` collection, `filterByTopic` filter, `sortEntriesByYear`/`sortEntriesByTopics` filters |
| `src/styles/tailwind.css` | Tailwind v4 config — `@plugin "@tailwindcss/typography"`, dark mode variant, diff token styles |

## Blog Post Context

**Three narrative threads** agreed with Niko:
1. **The hype test** — is OpenClaw actually useful or just noise?
2. **The security/privacy question** — Tailscale helps the network layer, but LLMs seeing personal data is unsolved (venice.ai, self-hosted LLM)
3. **The self-hosting spark** — this project reignited broader self-hosting interest

**Niko's stated motivations**:
- Mac mini collecting dust since MBP M5 purchase in December
- Curiosity about OpenClaw hype in RSS feeds
- Testing if security concerns are valid
- Learning Nix (AI did most of the config work, writing the blog is partly how he'll understand it)
- Mac Mini shortage gave a smugness factor

**Tone**: Practical, honest, personal. Not a tutorial. "Here's what actually happened."

**Important**: Niko writes the prose himself. The agent's role is outlining, structuring, reviewing drafts, and giving feedback. Do NOT write full paragraphs of blog content unless explicitly asked.

**Raw source material**: Available via SSH on Mac mini:
- `ssh mini "cat ~/.openclaw/workspace/memory/2026-02-07.md"`
- `ssh mini "cat ~/.openclaw/workspace/memory/2026-02-08.md"`
- `ssh mini "cat ~/.openclaw/workspace/MEMORY.md"`

## Resume Instructions

1. Run `bun start` in `/Users/nkdem/Developer/personal-site` (starts Eleventy + Tailwind watch)
2. Open `http://localhost:8081/blog/entry/first-post/` to see current state
3. Read `src/blog/entry/first-post.md` for current draft content
4. Continue helping Niko fill in the 4 sections — review his drafts, suggest improvements, keep it concise
5. When ready to publish: set `draft: false` in frontmatter, commit, push

## Warnings

- Tailwind v4 uses `@plugin` in CSS, not `plugins` in config. The `tailwind.config.js` is essentially unused (has a TODO saying so).
- The `{% terminal %}` shortcode is Nunjucks but Eleventy processes `.md` files with Liquid by default. It works because `{% %}` syntax overlaps, but be aware if adding Nunjucks-only features.
- `package.json` is missing `"type": "module"` — Node warns about ESM reparsing on every build. Non-breaking but noisy.
