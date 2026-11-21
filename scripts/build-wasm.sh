#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
goroot=$(go env GOROOT)
wasm_exec=""

for candidate in "$goroot/lib/wasm/wasm_exec.js" "$goroot/misc/wasm/wasm_exec.js"; do
	if [ -f "$candidate" ]; then
		wasm_exec=$candidate
		break
	fi
done

if [ -z "$wasm_exec" ]; then
	echo "Could not find wasm_exec.js in $goroot; install the full Go toolchain." >&2
	exit 1
fi

mkdir -p "$root/site"
cd "$root"
GOOS=js GOARCH=wasm go build -o site/strawberry.wasm ./cmd/wasm
cp "$wasm_exec" site/wasm_exec.js
