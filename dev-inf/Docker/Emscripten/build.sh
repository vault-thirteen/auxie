#!/bin/sh

# Exit on error.
set -e

img_hash_file="image_hash.txt"

docker build --file Emscripten.dockerfile .
docker images --no-trunc --format "{{.ID}}" | head -n 1 > $img_hash_file

img_hash=$(cat $img_hash_file)
echo "Image hash sum: $img_hash"
