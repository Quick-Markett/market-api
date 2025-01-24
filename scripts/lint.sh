set -e

GOLANGCI_LINT_VERSION="1.61.0"

CDPATH="" cd -- "$(dirname -- "$0")/.."
BIN="$(pwd -P)"/bin

mkdir -p "$BIN"

EXIT_CODE=0

fail() {
  echo "$@"
  EXIT_CODE=1
}

if ! "$BIN"/golangci-lint --version 2> /dev/null | grep -q "$GOLANGCI_LINT_VERSION"; then
  GOBIN="$BIN" go install "github.com/golangci/golangci-lint/cmd/golangci-lint@v$GOLANGCI_LINT_VERSION"
fi

MOD_DIRS="$(git ls-files '*go.mod' | xargs dirname | sort)"

for dir in $MOD_DIRS; do
  [ "$dir" = "example/newreposecretwithlibsodium" ] && continue
  echo linting "$dir"
  (
    cd "$dir"
    if [ -n "$GITHUB_ACTIONS" ]; then
      "$BIN"/golangci-lint run --path-prefix "$dir" --out-format colored-line-number
    else
      "$BIN"/golangci-lint run --path-prefix "$dir"
    fi
  ) || fail "failed linting $dir"
done

echo validating generated files
script/generate.sh --check || fail "failed validating generated files"

[ -z "$FAILED" ] || exit 1

exit "$EXIT_CODE"
