#!/bin/sh

day="$1"

outputfile=$day
files=$(find "day$day" -name '*.go')

echo $files

if [ -n "$files" ]; then
  go build -o "$outputfile" $files
  exit $?
else
  echo "No .go files found for day$day"
  exit 1
fi
