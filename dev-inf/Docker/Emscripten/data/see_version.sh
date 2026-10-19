#!/bin/bash

# This script prints the version of Emscripten SDK. The Docker container must 
# be started before using this script. This script must be called from the 
# "Emscripten" container.

# Exit on error.
set -e

# Apply Emscripten's paths.
echo PREPARING THE SDK ...
cd /home/emscripten/emsdk
source ./emsdk_env.sh
echo

# Print the version of the tools.
echo PRINTING VERSIONS.
lsb_release -a
git version
python3 --version
cmake --version
xz --version
emcc --check
echo
