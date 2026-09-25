# AltForge

A tiny Warcraft-inspired first-and-last-name generator for the alt you will definitely level this time. Built as a disposable playground for Lathe runs.

## Run

Requires Go 1.22 or newer. No database, API keys, npm, or build pipeline.

```sh
go run .
```

Open http://127.0.0.1:8088. To change the address, run `go run . -addr 127.0.0.1:9090`.

## Test

```sh
go test ./...
```

## How it works

Go's standard library serves embedded HTML templates and assets. The form uses HTMX to replace the result section with a server-rendered fragment. It also works as an ordinary GET form without JavaScript. Six ancestry styles each produce six distinct first-and-last-name combinations. These are invented fantasy names, not a guarantee of in-game naming rules or availability.

`names.go` holds the name pools and generator; `main.go` handles HTTP; `templates/page.html` contains the full page and results fragment; `static/style.css` styles the page. HTMX 2.0.11 is vendored in `static/htmx.min.js` with its license in `static/HTMX-LICENSE.txt`, from https://github.com/bigskysoftware/htmx. Documentation: https://htmx.org/docs/.

## Sample Lathe requests

- Add a shuffle option that chooses an ancestry at random and shows which one it selected.
- Let users choose between one, six, and twelve names per batch.
- Add Goblin names, with tests covering the new ancestry.
- Add a favorites list that persists across server restarts.
- Add a copy button to each name card with accessible confirmation feedback.

The repository starts with a clean local commit. `lathe plan` and `lathe implement` can use it directly. For `lathe build` and PR revision experiments, first create and configure a GitHub remote and push the baseline. No remote is created automatically.
