#!/bin/sh

# Exit on error.
set -e

docker build --file Emscripten.dockerfile .
