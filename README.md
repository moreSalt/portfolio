# Paul Rodriguez — Portfolio

A single-page, black-and-white portfolio generated from `content.md`.
One white page, monospace text, typed out section by section on load
like a terminal.

Built with Go, htmx, Tailwind CSS, and a tiny vanilla-JS typewriter.
Markdown is rendered with [goldmark](https://github.com/yuin/goldmark).

## Run

```bash
go run ./cmd/server
```

Then open http://localhost:8080 (set `PORT` to change the port).

## Deploy (Vercel)

The site runs on Vercel as a single Go serverless function
(`api/index.go`, built from the root `portfolio` package). A catch-all
rewrite in `vercel.json` routes `/`, `/frag/*`, and `/static/*` to it.

To deploy: push to GitHub and import the repo at
[vercel.com/new](https://vercel.com/new) — Vercel detects the Go
function automatically, no build settings needed. Or with the CLI:

```bash
vercel
```

## Content

Everything on the page comes from `content.md`:

- The `# heading` at the top becomes the intro.
- Each `## heading` starts a new section, fetched and typed in order.
- Standard Markdown works: paragraphs, **bold**, [links](...), lists.
- External links (http/https) automatically open in a new tab.

A `content.md` on disk overrides the one embedded in the binary, so you
can edit content and just restart the server — no rebuild needed.

## Styling (Tailwind)

Styles use the Tailwind v4 standalone CLI (no Node.js required). Download
the binary from the
[Tailwind releases page](https://github.com/tailwindlabs/tailwindcss/releases)
as `tailwindcss`, then:

```bash
./tailwindcss -i input.css -o static/style.css --minify   # add --watch while editing
```

`input.css` scans `templates/` for class names. Rebuild the Go binary
afterwards (templates and static files are embedded).

## How it works

- `GET /` serves a tiny HTML shell.
- At startup, `content.md` is split at each `##` heading; every section
  becomes an htmx fragment endpoint at `/frag/{slug}` (slug derived from
  the heading), chained in document order.
- The shell fetches the first fragment on load; the typewriter types it
  with a human-ish cadence (jitter, punctuation pauses), fires a `typed`
  event, and the next fragment is fetched — chaining down the page while
  auto-scrolling.
- The `skip »` button (bottom right) instantly finishes all typing and
  hides itself. Honors `prefers-reduced-motion` (everything renders
  instantly).
- With JS disabled, a `<noscript>` block renders the full page statically.
