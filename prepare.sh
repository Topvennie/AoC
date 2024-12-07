#!/bin/sh

if [ -z "$(command -v go)" ]; then
    echo "Install Golang v1.23.3"
    exit 1
fi

v=`go version | { read _ _ v _; echo ${v#go}; }`
if [ "$v" != "1.23.3" ]; then
    echo "Golang version is not 1.23.3. Let's hope for the best."
fi

year="2024"
day="$1"

outputfile="${year}_${day}"
files=$(find "$year/day$day" -name '*.go')

if [ -n "$files" ]; then
    go build -o "$outputfile" $files
    exit $?
else
    echo "No .go files found for $year/day$day"
    exit 1
fi
