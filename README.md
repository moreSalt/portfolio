# Paul Rodriguez — Portfolio

A single-page, black-and-white portfolio generated from `content.md`.
One white page, monospace text, fully static — no JavaScript.

Built with Go, Tailwind CSS, and Markdown rendered with
[goldmark](https://github.com/yuin/goldmark).

## Run

```bash
go build -o portfolio .
./portfolio
```

Then open http://localhost:8080 (set `PORT` to change the port).

## Deploy (Vercel)

Push to GitHub and import the repo at
[vercel.com/new](https://vercel.com/new). Vercel detects the Go server
(`main.go` at the repo root), builds it, and routes all traffic to it —
no configuration files needed. The server reads the port from `PORT`,
which Vercel sets automatically.

## Content

Everything on the page comes from `content.md`:

- The `# heading` at the top becomes the intro.
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

- At startup, `content.md` is rendered to HTML and `style.css` is read in
  one pass.
- `GET /` serves the full page inside a small HTML shell with the CSS
  inlined (no render-blocking requests); `/static/*` serves the favicon
  with a year-long immutable cache header.
- Text responses are gzip-compressed, and no JavaScript is shipped at all.
