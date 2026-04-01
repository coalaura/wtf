#!/bin/bash
set -e

OS=$(uname -s | tr 'A-Z' 'a-z')

ARCH=$(uname -m)
case "$ARCH" in
	x86_64)
		ARCH=amd64
		;;
	aarch64|arm64)
		ARCH=arm64
		;;
	*)
		echo "Unsupported architecture: $ARCH" >&2
		exit 1
		;;
esac

echo "Resolving latest version of wtf..."

VERSION=$(curl -sL https://api.github.com/repos/coalaura/wtf/releases/latest | grep -Po '"tag_name": *"\K.*?(?=")')

if ! printf '%s\n' "$VERSION" | grep -Eq '^v[0-9]+\.[0-9]+\.[0-9]+$'; then
	echo "Error: '$VERSION' is not in vMAJOR.MINOR.PATCH format" >&2
	exit 1
fi

rm -f /tmp/wtf

BIN="wtf_${VERSION}_${OS}_${ARCH}"
URL="https://github.com/coalaura/wtf/releases/download/${VERSION}/${BIN}"

echo "Downloading ${BIN}..."

if ! curl -sL "$URL" -o /tmp/wtf; then
	echo "Error: failed to download $URL" >&2
	exit 1
fi

trap 'rm -f /tmp/wtf' EXIT

chmod +x /tmp/wtf

echo "Installing to /usr/local/bin/wtf requires sudo"

if ! sudo install -m755 /tmp/wtf /usr/local/bin/wtf; then
	echo "Error: install failed" >&2
	exit 1
fi

echo "wtf $VERSION installed to /usr/local/bin/wtf"

if [ -f "$HOME/.bashrc" ] && ! grep -qF 'wtf --completion bash' "$HOME/.bashrc" 2>/dev/null; then
	echo 'eval "$(wtf --completion bash)"' >> "$HOME/.bashrc"

	echo "Added bash completion to ~/.bashrc"
fi
