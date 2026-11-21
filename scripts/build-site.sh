#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)

"$root/scripts/build-wasm.sh"

if ! command -v jekyll >/dev/null 2>&1; then
	echo "Jekyll is required to build the documentation site." >&2
	echo "Install it with: gem install jekyll" >&2
	exit 1
fi

jekyll build --source "$root/docs" --destination "$root/site/docs"
