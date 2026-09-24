#!/bin/sh

# Exit on error.
set -e

img_hash_file="image_hash.txt"
img_hash=$(cat $img_hash_file)

docker run -it \
    --name Emscripten \
    --mount type=bind,source=/home/username/Docker/Emscripten/data,destination=/home/data \
    $img_hash
