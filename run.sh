#!/bin/sh

day="$(printf '%02d' ${1##0})"
part="$2"
input="$3"

year="2025"

./"$year"/"$day" -p "$part" -i "$input"
