#!/bin/sh

# Exit on error.
set -e

docker run -it \
    --name Emscripten \
    --mount type=bind,source=/home/username/Docker/Emscripten/data,destination=/home/data \
    sha256:<sha256_sum_of_the_image>
