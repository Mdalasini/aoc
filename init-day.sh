#!/usr/bin/env bash
set -euo pipefail

usage() {
  echo "Usage: $0 <year> <day>" >&2
  echo "Example: $0 2020 1  # creates 20201/" >&2
  exit 1
}

if [[ $# -ne 2 ]]; then
  usage
fi

year="$1"
day="$2"

if ! [[ "$year" =~ ^[0-9]{4}$ ]]; then
  echo "Error: year must be a 4-digit number (got '$year')" >&2
  exit 1
fi

if ! [[ "$day" =~ ^[0-9]+$ ]] || (( day < 1 || day > 25 )); then
  echo "Error: day must be between 1 and 25 (got '$day')" >&2
  exit 1
fi

dir="${year}${day}"

if [[ -e "$dir" ]]; then
  echo "Error: '$dir' already exists" >&2
  exit 1
fi

mkdir "$dir"
touch "$dir/sample.txt" "$dir/input.txt"

cat > "$dir/main.go" <<EOF
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("sample.txt")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	_ = scanner
}
EOF

echo "Created $dir/ with main.go, sample.txt, and input.txt"
