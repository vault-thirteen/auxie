#!/bin/sh

# Exit on error.
set -e

docker run --name Emscripten -it --mount type=bind,source=/home/username/Docker/Emscripten/data,destination=/home/data sha256:b15780ebf6443e3eec8b999a2fdeeb16816ad203eeabd8b6dd12f5c3e8f6be82
