#!/bin/sh

pattern="??"

for binary in $pattern; do
  if [ -f "$binary" ]; then
    echo "Removing binary: $binary"
    rm "$binary"
  fi
done

echo "Cleanup complete!"
