#!/usr/bin/env bash
set -euo pipefail

usage() {
  echo "Usage: $0 <year> <day>" >&2
  echo "Example: $0 2020 1  # creates 20201/" >&2
  exit 1
}

fetch_aoc_input() {
  local year="$1"
  local day="$2"
  local outfile="$3"

  local session="${AOC_SESSION:-}"
  if [[ -z "$session" && -f ".aoc_session" ]]; then
    session="$(tr -d '[:space:]' < .aoc_session)"
  fi

  if [[ -z "$session" ]]; then
    echo "Warning: set AOC_SESSION or create .aoc_session to fetch puzzle input" >&2
    return 0
  fi

  local url="https://adventofcode.com/${year}/day/${day}/input"
  if ! curl -sf -H "Cookie: session=${session}" "$url" -o "$outfile"; then
    echo "Error: failed to fetch input from ${url}" >&2
    rm -f "$outfile"
    return 1
  fi

  if grep -q "Please log in to get your puzzle input" "$outfile"; then
    echo "Error: invalid session cookie; got login page instead of puzzle input" >&2
    rm -f "$outfile"
    return 1
  fi

  echo "Fetched puzzle input to ${outfile}"
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
touch "$dir/sample.txt"

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

fetch_aoc_input "$year" "$day" "$dir/input.txt" || touch "$dir/input.txt"

echo "Created $dir/ with main.go, sample.txt, and input.txt"
