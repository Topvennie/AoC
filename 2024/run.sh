#!/bin/sh

day="$1"
part="$2"
input="$3"

binary=$day
flags="-p $part -i $input"

if [ -x "$binary" ]; then
  ./"$binary" $flags
else
  echo "Error: Binary '$binary' not found or not executable."
  exit 1
fi
