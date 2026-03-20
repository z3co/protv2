#!/bin/sh

OS=$(uname -s | tr '[:upper:]' '[:lower:]')
ARCH=$(uname -m)
VERSION="1.0.1"

case $ARCH in
	x86_64) ARCH="amd64" ;;
	aarch64|arm64) ARCH="arm64" ;;
	*) echo "Unsupported arch $ARCH" ;;
esac

TMP_DIR=$(mktemp -d)
cd $TMP_DIR
echo "Downloading tarball"
TARBALL="prot-${VERSION}-${OS}-${ARCH}.tar.gz"
URL="https://github.com/z3co/protv2/releases/download/v${VERSION}/${TARBALL}"
CHECKSUM="https://github.com/z3co/protv2/releases/download/v${VERSION}/checksums.txt"
echo "$URL"
curl -LO "$URL"
curl -LO "$CHECKSUM"

echo "Verifying checksums..."
if command -v sha256sum >/dev/null 2>&1; then
	sha256sum --check --ignore-missing checksums.txt
else
	shasum -a 256 --check --ignore-missing checksums.txt
fi

echo "Installing program"
tar -xzf "${TARBALL}"
sudo mv prot /usr/local/bin/prot
sudo chmod +x /usr/local/bin/prot

cd /
rm -rf "${TMP_DIR}"
echo "Done"
