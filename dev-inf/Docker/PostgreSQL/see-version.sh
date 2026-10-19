#!/bin/bash

# Exit on error.
set -e

docker inspect postgresql | grep VERSION
