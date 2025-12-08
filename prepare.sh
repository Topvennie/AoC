#!/bin/sh

day="$(printf '%02d' ${1##0})"

cd 2025 || 1

./prepare.sh $day
