# Strawberry

Strawberry is a small interpreter written in Go. It is an experiment in building a programming language and is still under active development.

## Try the interpreter

Run the command-line REPL:

```sh
go run .
```

Run the tests:

```sh
go test ./...
```

The evaluator currently supports integer arithmetic, comparisons, booleans, and prefix operators. See the [language notes](docs/language.md) for what is and is not implemented yet.

## Website and playground

The static playground compiles Strawberry to WebAssembly and runs programs locally in your browser. Build it together with the separate Jekyll documentation site:

```sh
./scripts/build-site.sh
python3 -m http.server 8000 --directory site
```

Open <http://localhost:8000>. Requirements: Go 1.22 or newer with its WebAssembly runtime files, Jekyll, and Python 3 (or another static HTTP server).

Build only the WebAssembly interpreter with `./scripts/build-wasm.sh`. The shared API used by both the REPL and playground is in [`interpreter/`](interpreter/).
