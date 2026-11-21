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

## Browser playground and docs

The browser playground has a REPL-style workspace: choose an example, edit a program, and run it beside its output. Strawberry is compiled to WebAssembly, so code runs locally in your browser and is not sent to a server.

The documentation is a separate Jekyll site. Install Jekyll if needed:

```sh
gem install jekyll
```

Build both the playground and documentation, then serve the generated site:

```sh
./scripts/build-site.sh
python3 -m http.server 8000 --directory site
```

Open <http://localhost:8000>. You’ll need Go 1.22 or newer with its WebAssembly runtime files, Jekyll, and Python 3 (or another static HTTP server).

To build only the WebAssembly interpreter, run `./scripts/build-wasm.sh`. The shared API used by both the command-line REPL and playground is in [`interpreter/`](interpreter/).

For the language’s current capabilities and limitations, see the [language notes](docs/language.md).

## Deploy on Vercel

Import the repository with its root directory as the project root and use the **Other** framework preset. [`vercel.json`](vercel.json) installs Jekyll, downloads a checksum-verified Go toolchain for the build, and publishes `site/`. Push changes to your deployment branch to trigger a new build.
