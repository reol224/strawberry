#!/bin/sh
set -eu

root=$(CDPATH= cd -- "$(dirname -- "$0")/.." && pwd)
version=1.25.14

case "$(uname -s)-$(uname -m)" in
	Linux-x86_64)
		arch=amd64
		checksum=a21ae5633a269bcd7e90cf767e48225633795e99d831742cbf3397064fee7712
		;;
	Linux-aarch64)
		arch=arm64
		checksum=9bf234ea70ffec9347fdf6b22ce4add51717d3386a38a441e8c8743fceb5eaee
		;;
	*)
		echo "Unsupported Vercel build platform: $(uname -s)-$(uname -m)" >&2
		exit 1
		;;
esac

toolchain_dir=$(mktemp -d)
archive="$toolchain_dir/go.tar.gz"
trap 'rm -f "$archive"' EXIT

curl --fail --location --silent --show-error \
	"https://go.dev/dl/go${version}.linux-${arch}.tar.gz" -o "$archive"
printf '%s  %s\n' "$checksum" "$archive" | sha256sum --check
tar -xzf "$archive" -C "$toolchain_dir"
PATH="$toolchain_dir/go/bin:$PATH"
export PATH

"$root/scripts/build-site.sh"
