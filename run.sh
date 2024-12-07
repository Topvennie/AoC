#!/bin/sh

year="2024"
day="$1"
part="$2"
input="$3"

binary="${year}_${day}"
flags="-p $part -i $input"

if [ -x "$binary" ]; then
    ./"$binary" $flags  # Separate the binary name and arguments
else
    echo "Error: Binary '$binary' not found or not executable."
    exit 1
fi
