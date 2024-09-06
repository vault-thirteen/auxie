#!/bin/sh

# Exit on error.
set -e

docker run --name Emscripten -it --mount type=bind,source=/home/xxx/Docker/Emscripten/data,destination=/home/data sha256:xxx
