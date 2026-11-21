---
layout: default
title: Build and run
---

## Requirements

- Go 1.22 or newer, with the `js/wasm` target and `wasm_exec.js` included in the installation.
- Jekyll to build the documentation pages.
- Python 3 (or another static HTTP server) to preview the built site.

## Run the interpreter

From the repository root:

```sh
go run .
```

## Build and preview the website

```sh
./scripts/build-site.sh
python3 -m http.server 8000 --directory site
```

Then open `http://localhost:8000`. The playground is built to WebAssembly and runs in the browser; it does not send source code to a server.

## Test the Go code

```sh
go test ./...
```

The interpreter API used by both the command-line REPL and browser bridge is in `interpreter/`.
