#!/usr/bin/env bash
# Fetch Advent of Code puzzle input
# Usage: ./fetch-input.sh <year> <day>
# Requires: AOC_SESSION env var with your adventofcode.com session cookie

set -euo pipefail

if [ -z "${AOC_SESSION:-}" ]; then
  echo "Error: AOC_SESSION env var not set"
  echo "Get it from: adventofcode.com → Dev Tools → Cookies → session"
  exit 1
fi

if [ $# -ne 2 ]; then
  echo "Usage: ./fetch-input.sh <year> <day>"
  echo "Example: ./fetch-input.sh 2024 9"
  exit 1
fi

YEAR=$1
DAY=$2
DAY_PADDED=$(printf '%02d' "$DAY")
DIR="${YEAR}/${DAY_PADDED}"
OUTPUT="${DIR}/input.txt"

if [ ! -d "$DIR" ]; then
  echo "Error: directory $DIR does not exist"
  exit 1
fi

curl -s -b "session=${AOC_SESSION}" \
  "https://adventofcode.com/${YEAR}/day/${DAY}/input" \
  -o "$OUTPUT"

LINES=$(wc -l < "$OUTPUT" | tr -d ' ')
echo "Saved ${LINES} lines to ${OUTPUT}"
